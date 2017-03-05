package maze_test

import (
	"github.com/bigmikef/maze"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Maze", func() {
	var (
		simple      maze.Room
		oneDoor     maze.Room
		twoDoor     maze.Room
		cycleExit   maze.Room
		cycleNoExit maze.Room
	)

	BeforeEach(func() {
		simple = maze.Room{
			Exit: true,
		}

		oneDoor = maze.Room{
			West: &simple,
		}

		twoDoor = maze.Room{
			North: &oneDoor,
		}

		var cycleExitPrime = maze.Room{
			North: &cycleExit,
			West:  &twoDoor,
		}

		cycleExit = maze.Room{
			South: &cycleExitPrime,
		}

		var cycleNoExitTwo = maze.Room{
			North: &cycleNoExit,
		}

		var cycleNoExitOne = maze.Room{
			West: &cycleNoExitTwo,
		}

		cycleNoExit = maze.Room{
			South: &cycleNoExitOne,
		}
	})

	Describe("A maze of one room with an exit", func() {
		It("should find exit", func() {
			Expect(maze.FindExit(&simple)).To(Equal(true))
		})
	})

	Describe("A maze of 2 rooms with exit in other room", func() {
		It("should find exit", func() {
			Expect(maze.FindExit(&oneDoor)).To(Equal(true))
		})
	})

	Describe("A maze of 3 rooms with exit in other room", func() {
		It("should find exit", func() {
			Expect(maze.FindExit(&twoDoor)).To(Equal(true))
		})
	})

	Describe("A maze with a cycle", func() {
		It("should find exit", func() {
			Expect(maze.FindExit(&cycleExit)).To(Equal(true))
		})
	})

	Describe("A maze with a cycle and no exit", func() {
		It("should not find exit", func() {
			Expect(maze.FindExit(&cycleNoExit)).To(Equal(false))
		})
	})
})
