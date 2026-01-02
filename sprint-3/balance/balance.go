package balance

type Stack struct {
	stack []string
	top   string
	len   uint64
}

func (s *Stack) New() *Stack {
	s.stack = make([]string, 0)
	s.top = ""
	s.len = 0
	return s
}

func (s *Stack) Push(e string) uint64 {
	s.stack = append(s.stack, e)
	s.top = e
	s.len++
	return s.len
}

func (s *Stack) Pop() (string, bool) {
	if s.len == 0 {
		return "", false
	}
	p := s.top
	s.stack = s.stack[:s.len-1]
	s.len--
	if s.len == 0 {
		s.top = ""
	} else {
		s.top = s.stack[s.len-1]
	}
	return p, true
}

func (s *Stack) Top() string {
	return s.top
}

func (s *Stack) Len() uint64 {
	return s.len
}

func Balance(s string) bool {
	var st Stack
	st.New()

	for _, r := range s {
		switch string(r) {

		case "[", "(", "{":
			st.Push(string(r))

		case "]":
			if st.Len() == 0 {
				return false
			}
			if st.Top() != "[" {
				return false
			}
			st.Pop()
		case ")":
			if st.Len() == 0 {
				return false
			}
			if st.Top() != "(" {
				return false
			}
			st.Pop()
		case "}":
			if st.Len() == 0 {
				return false
			}
			if st.Top() != "{" {
				return false
			}
			st.Pop()
		}

	}

	return st.len == 0
}
