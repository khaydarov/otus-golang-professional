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
		nextStageStream := make(Bi)
		go func(stage Stage, in In, nextStageStream Bi) {
			defer close(nextStageStream)

			for v := range stage(in) {
				select {
				case nextStageStream <- v:
				case <-done:
					return
				}
			}
		}(stage, lastStageStream, nextStageStream)
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
