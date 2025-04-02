package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func stageWithDone(in In, done In, stage Stage) Out {
	if in == nil {
		return nil
	}

	out := make(Bi)
	go func() {
		defer close(out)
		stageStream := stage(in)
		for v := range stageStream {
			select {
			case <-done:
				return
			case out <- v:
			}
		}
	}()
	return out
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if len(stages) == 0 {
		return in
	}

	result := in
	for _, stage := range stages {
		if stage == nil {
			continue
		}
		result = stageWithDone(result, done, stage)
	}
	return result
}
