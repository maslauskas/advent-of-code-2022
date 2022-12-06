package day06

import "testing"

func TestExample(t *testing.T) {
	t.Run("will find packets", func(t *testing.T) {
		dataset := map[string]int{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    7,
			"bvwbjplbgvbhsrlpgdmjqwftvncz":      5,
			"nppdvjthqldpwncqszvftbrmjlhg":      6,
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 10,
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  11,
		}

		for packet, want := range dataset {
			got := FindMarker(packet)

			if got != want {
				t.Errorf("expected packet marker to be %d, got %d with packet %q", want, got, packet)
			}
		}
	})

	t.Run("will detect messages", func(t *testing.T) {
		dataset := map[string]int{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb":    19,
			"bvwbjplbgvbhsrlpgdmjqwftvncz":      23,
			"nppdvjthqldpwncqszvftbrmjlhg":      23,
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg": 29,
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw":  26,
		}

		for packet, want := range dataset {
			got := FindMessage(packet)

			if got != want {
				t.Errorf("expected packet message to be %d, got %d with packet %q", want, got, packet)
			}
		}
	})
}
