package ast

type CollectFnc func(filePath string) error

func All(collectors ...CollectFnc) CollectFnc {
	return func(filePath string) error {
		for _, c := range collectors {
			if err := c(filePath); err != nil {
				return err
			}
		}

		return nil
	}
}