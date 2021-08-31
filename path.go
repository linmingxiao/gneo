package gneo

func cleanPath(p string) string{
	const stackBufSize = 128
	if p==""{
		return "/"
	}

	buf := make([]byte, 0, stackBufSize)
	n:= len(p)
	r := 1
	w := 1

	if p[0] != '/'{
		r=0
		if n+1 > stackBufSize{
			buf = make([]byte, n+1)
		} else{
			buf = buf[:n+1]
		}
		buf[0] = '/'
	}

	trailing := n > 1 && p[n-1] == '/'

	for r < n {
		switch {
		case p[r] == '/':
			// empty path element, trailing slash is added after the end
			r++

		case p[r] == '.' && r+1 == n:
			trailing = true
			r++

		case p[r] == '.' && p[r+1] == '/':
			// . element
			r += 2

		case p[r] == '.' && p[r+1] == '.' && (r+2 == n || p[r+2] == '/'):
			// .. element: remove to last /
			r += 3

			if w > 1 {
				// can backtrack
				w--

				if len(buf) == 0 {
					for w > 1 && p[w] != '/' {
						w--
					}
				} else {
					for w > 1 && buf[w] != '/' {
						w--
					}
				}
			}

		default:
			// Real path element.
			// Add slash if needed
			if w > 1 {
				bufApp(&buf, p, w, '/')
				w++
			}

			// Copy element
			for r < n && p[r] != '/' {
				bufApp(&buf, p, w, p[r])
				w++
				r++
			}
		}
	}

	// Re-append trailing slash
	if trailing && w > 1 {
		bufApp(&buf, p, w, '/')
		w++
	}

	// If the original string was not modified (or only shortened at the end),
	// return the respective substring of the original string.
	// Otherwise return a new string from the buffer.
	if len(buf) == 0 {
		return p[:w]
	}
	return string(buf[:w])
}

func bufApp(buf *[]byte, s string, w int, c byte){
	b:= *buf
	if len(b) == 0{
		if s[w] == c{
			return
		}

		length := len(s)
		if length > cap(b){
			*buf = make([]byte, length)
		} else {
			*buf = (*buf)[:length]
		}
		b = *buf
		copy(b, s[:w])
	}
	b[w] = c
}
