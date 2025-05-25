# Robot API Server

A simple REST API server for controlling and managing robots in a virtual environment. This project was created as part of the MKSS2 course.

## Features

- Get robot status (position, energy, inventory)
- Move robots in different directions
- Pickup and put down items
- Update robot state
- Track robot actions
- Robot combat system
- API documentation with Swagger

## Setup and Installation

### Installation

1. Clone the repository:

   ```
   git clone <repository-url>
   cd aufgabe-1
   ```

2. Install dependencies:

   ```
   go mod download
   ```

3. Run the server:
   ```
   go run .
   ```

The server will start on port 8080.

## Testing with Postman

### Import the Collection

1. Download the Postman collection file from the project directory:

   - `Robot_API.postman_collection.json`

2. Open Postman and click on "Import" in the top left corner

3. Drag and drop the collection file or click "Upload Files" to select it

4. The Robot API collection will be imported with all endpoints ready to test

### Using the Collection

The collection includes all available endpoints with pre-configured:

- Request methods (GET, POST, PATCH)
- URL parameters
- Example request bodies for POST/PATCH requests

To test the API:

1. Start the server: `go run .`
2. Open the imported collection in Postman
3. Select an endpoint and click "Send"
4. View the response in the response panel

### Collection Environment Variables

The collection uses the following variables:

- `baseUrl`: Set to `http://localhost:8080` by default
- `robotId`: Set to `robot1` by default
- `itemId`: Set to `item1` by default

You can modify these variables in the Postman environment settings.

## API Endpoints

### Get Robot Status

```
GET /robot/:id/status
```

Returns the current status of a robot including position, energy, and inventory.

### Move Robot

```
POST /robot/:id/move
```

Move a robot in a specific direction.

Request body:

```json
{
	"direction": "up" // "up", "down", "left", "right"
}
```

### Pickup Item

```
POST /robot/:id/pickup/:itemId
```

Makes a robot pick up an item and add it to its inventory.

### Put Down Item

```
POST /robot/:id/putdown/:itemId
```

Makes a robot put down an item from its inventory.

### Update Robot State

```
PATCH /robot/:id/state
```

Updates robot's state (energy and/or position).

Request body:

```json
{
	"energy": 100,
	"position": {
		"x": 10,
		"y": 20
	}
}
```

### Get Robot Actions

```
GET /robot/:id/actions
```

Returns a history of actions performed by the robot.

### Attack Robot

```
POST /robot/:id/attack/:targetId
```

Makes one robot attack another, reducing the target's energy.

## Examples

### Get Robot Status

```
curl -X GET http://localhost:8080/robot/robot1/status
```

### Move Robot

```
curl -X POST http://localhost:8080/robot/robot1/move -H "Content-Type: application/json" -d '{"direction": "up"}'
```

### Pickup Item

```
curl -X POST http://localhost:8080/robot/robot1/pickup/item1
```

## Initial Data

The server starts with two robots:

- robot1 at position (0,0) facing north
- robot2 at position (10,10) facing south

And three items: item1, item2, item3
