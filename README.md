# Parking Lot - LLD

## The Parking Lot Problem Unveiled
Imagine a Saturday afternoon at a popular shopping mall. Vehicles of all shapes and sizes flood the parking lot, creating a chaotic scene. Cars zigzag, motorcycles weave, and trucks juggle for space. The challenges of efficiently allocating spots and avoiding traffic jams become evident.

In the world of software, the parking lot problem is a classic conundrum that mirrors this real-world scenario. It involves designing a system that can seamlessly manage the allocation and deallocation of parking spots, accommodating vehicles of different types and sizes while maintaining order.

## Designing a Modern Parking Lot Solution
To tackle the parking lot problem, we need to break down the solution into its essential components:

- **Vehicles**: Different types of vehicles such as cars, motorcycles, and trucks, each with unique attributes.
- **Parking Spots**: Individual parking spots that need to be allocated and deallocated efficiently.
- **Parking Lot Manager**: The brain behind the operation, responsible for coordinating vehicle arrivals and departures.

But let’s add a twist: imagine the parking lot has multiple gates. This feature introduces a new layer of complexity. Vehicles now arrive from various directions, making the simulation more realistic and challenging.

## Building with Go: Concurrency and Synchronization
Enter Go, the programming language known for its robust support for concurrency. Concurrency allows multiple tasks to run concurrently, mimicking the simultaneous actions of vehicles in our parking lot.

One crucial concept in concurrent programming is **mutexes** (short for mutual exclusion). A mutex acts as a gatekeeper, allowing only one thread to access a shared resource at a time. This is crucial to prevent data race conditions where multiple threads access the same resource simultaneously, leading to unpredictable behavior.

Here’s a snippet showcasing how mutexes work in Go:

```go
import "sync"
// Create a mutex
var mu sync.Mutex
// Lock the mutex before accessing shared resource
mu.Lock()
// Access shared resource
mu.Unlock()
```

## Implementing the Solution
Let’s dive into the actual implementation of our parking lot solution using Go. We’ll focus on the parking and unparking operations, along with the synchronization of multiple gates.

```go
// …
type ParkingLot struct {
  spots []ParkingSpot
  mu sync.Mutex
}
func (p *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
  p.mu.Lock()
  defer p.mu.Unlock()
  // Implement parking logic here
}
func (p *ParkingLot) UnparkVehicle(vehicle Vehicle) bool {
  p.mu.Lock()
  defer p.mu.Unlock()
  // Implement unparking logic here
}
// …
```

## Feature Enhancement: The Magic of Multiple Gates
Our parking lot solution gains an exciting twist with the addition of multiple gates. This feature mimics real-world scenarios where vehicles enter from different directions. To simulate this, we can create concurrent goroutines that represent vehicles arriving from various gates.

Here’s a simplified snippet demonstrating this concept:

```go
  // …
  var wg sync.WaitGroup
  numVehicles := 10

  wg.Add(numVehicles)

  for i := 0; i < numVehicles; i++ {
  go func(vehicleIndex int) {
    defer wg.Done()
    // Simulate vehicle arriving from gate
    }(i)
  }
  wg.Wait()
  // …
```

Check here for the complete code.
