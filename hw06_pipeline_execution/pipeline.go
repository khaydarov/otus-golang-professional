package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	pipelineOut := make(Bi)
	lastStageStream := in
	for _, stage := range stages {
		// Creating intermediateStream to translate from global "in" channel to intermediary
		intermediary := make(Bi)
		go func(intermediary Bi, in In) {
			defer close(intermediary)

			for v := range in {
				select {
				case intermediary <- v:
				case <-done:
					return
				}
			}
		}(intermediary, lastStageStream)

		// Creating nextStageStream to translate current iteration stage channel data to the next one
		nextStageStream := make(Bi)
		go func(stage Stage, intermediary In, nextStageStream Bi) {
			defer close(nextStageStream)
			for v := range stage(intermediary) {
				select {
				case nextStageStream <- v:
				case <-done:
					return
				}
			}
		}(stage, intermediary, nextStageStream)
		lastStageStream = nextStageStream
	}

	go func() {
		defer close(pipelineOut)
		for {
			select {
			case <-done:
				return
			case v, ok := <-lastStageStream:
				if !ok {
					return
				}
				pipelineOut <- v
			}
		}
	}()

	return pipelineOut
}
