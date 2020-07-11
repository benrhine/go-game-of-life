# go-game-of-life
Implementation of Conway's game of life in Go

This is a implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway's_Game_of_Life).
Given a grid of both living and dead cells, cells will die and come to life as the game progresses through generations
based on some basic rules:

1. Any live cell with less than two live neighbours dies, as if caused by under-population.
2. Any live cell with two or three live neighbours lives on to the next generation.
3. Any live cell with greater than three live neighbours dies, as if by overcrowding.
4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

The game can be seeded with randomly generated cells of a given size, or load an initial cell from any text file.

Cell grid example:

```
......O.
OOO...O.
......O.
........
...OO...
...OO...
```


The nature of these rules creates scenarios where all cells may die (extinction) and where no cells can change their
state (utopia).  In either of these scenarios, the game will stop.

Note: All gradlew commands below are suggested to be executed with a clean to ensure you get a good test run or build.


### Running the tests

### Building and Running the Application

The application is packaged into a single executable file using the Go build tool.

```
go build cmd/go-game-of-life.go
```

From your projected directory, to execute

```
./go-game-of-life
```

### Testing and Building at the same time

If you would like to test and build the application at the same time.

```
???
```

Once packaged, a single executable file will be generated in your project root directory.

```
./go-game-of-life
```

There are a few arguments that can be specified:

* **-width <size>** : Set the width of the grid
* **-hieght <size>** : Set the hieght of the grid

```
./go-game-of-life -width=100 -height=100
```

Note: To see the Groovy of Life written in Groovy see my implementation [here](https://github.com/benrhine/GroovyGameOfLife/blob/master/README.md)

Note: To see the Java of Life written in Java 8 (which is based on this Groovy implementation) see my implementation [here](https://github.com/benrhine/JavaGameOfLife/blob/master/README.md)

