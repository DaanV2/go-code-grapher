package graphs

import "io"



type BaseWriter interface {
	io.Closer
	Start() error
	Finish() error
}

type StringCloserWriter interface {
	io.Closer
	io.StringWriter
}

type StringWriter struct {
	Writer StringCloserWriter
}

func (s *StringWriter) Write(content string) error {
	_, err := s.Writer.WriteString(content)

	return err
}

func (s *StringWriter) Close() error {
	return s.Writer.Close()
}

func (s *StringWriter) WriteLine(content string) error {
	return s.Write(content + "\n")
}

func (s *StringWriter) WriteLines(lines ...string) error {
	for _, line := range lines {
		if err := s.WriteLine(line); err != nil {
			return err
		}
	}

	return nil
}
