package time_needed_to_buy_tickets

// Time: O(n), Space: O(n)
func timeRequiredToBuyQueue(tickets []int, k int) (steps int) {
	if len(tickets) == 0 || tickets[k] == 0 {
		return
	}

	for {
		steps++
		tickets[0]--

		if k == 0 && tickets[0] == 0 {
			return
		}

		// Rotate the queue
		if tickets[0] == 0 {
			tickets = tickets[1:]
		} else {
			tickets = append(tickets[1:], tickets[0])
		}

		// Update k index
		if k == 0 {
			k = len(tickets) - 1
		} else {
			k--
		}
	}
}

// Time: O(n), Space: O(1)
func timeRequiredToBuyPointer(tickets []int, k int) (steps int) {
	if len(tickets) == 0 || tickets[k] == 0 {
		return
	}

	for { // 此处采用无限循环后，编译器不会要求之后必须有return语句
		for i := range tickets {
			if tickets[i] > 0 {
				tickets[i]--
				steps++
				if i == k && tickets[i] == 0 {
					return
				}
			}
		}
	}
}

// Time: O(n), Space: O(1)
func timeRequiredToBuyMath(tickets []int, k int) (steps int) {
	if len(tickets) == 0 || tickets[k] == 0 {
		return
	}

	for i, ticket := range tickets {
		if i <= k {
			// Each person before the target can buy at most tickets[k] tickets
			steps += min(ticket, tickets[k])
		} else {
			// Each person after the target can buy at most tickets[k]-1 tickets
			steps += min(ticket, tickets[k]-1)
		}
	}

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
