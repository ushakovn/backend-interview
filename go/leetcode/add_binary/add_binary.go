package add_binary

func addBinary(a string, b string) string {
	var (
		add bool
		num rune
		cur int
		sum []rune
	)
	ar := []rune(a)
	br := []rune(b)
	var (
		i int
		j int
	)
	for i, j = len(ar)-1, len(br)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {

		an := mustInt(ar[i])
		bn := mustInt(br[j])

		if cur = an + bn; add {
			cur += 1
		}
		switch cur {
		case 0:
			num = '0'
			add = false
		case 1:
			num = '1'
			add = false
		case 2:
			num = '0'
			add = true
		case 3:
			num = '1'
			add = true
		}
		sum = append(sum, num)
	}
	var (
		k  = -1
		mr []rune
	)
	if i >= 0 {
		k = i
		mr = ar
	} else if j >= 0 {
		k = j
		mr = br
	}
	for ; k >= 0; k-- {
		mn := mustInt(mr[k])

		if add {
			mn += 1
		}
		switch mn {
		case 0:
			num = '0'
			add = false
		case 1:
			num = '1'
			add = false
		case 2:
			num = '0'
			add = true
		}
		sum = append(sum, num)
	}
	if add {
		sum = append(sum, '1')
	}
	out := make([]rune, 0, len(sum))

	for k := len(sum) - 1; k >= 0; k-- {
		out = append(out, sum[k])
	}
	return string(out)
}

func mustInt(r rune) int {
	return int(r - '0')
}
