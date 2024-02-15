package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/thanhyarn/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://thanhyarn:123@atlascluster.bfqeyvn.mongodb.net/Ticket?retryWrites=true&w=majority"
const dbName = "TicketDB"
const colName = "Ticket"

//MOST IMPORTANT
var collection *mongo.Collection

// connect with monogoDB

func init() {
	// This line creates a MongoDB client options instance by calling the optons.Client() function, which returns a new options builder
	// Then it sets the connection string using the ApplyURL() method, passing the connectionString varuable

	clientOption := options.Client().ApplyURI(connectionString)

	// This line establishes a connection to the MOngoDB context.Context and a client options instance á parameters.
	// The context.TODO() function returns a non0-nil, empty context. It return s a *mongo.Client object representing the connection to the MongDB server and an error, if any.

	client, err := mongo.Connect(context.TODO(), clientOption)

	// This line checks if there was an error during the connection process.
	// If an error occurred , it logs the error using log.Fatal() and terminates the program

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	// This line sets the cglobal collection variable to refer to a specified database('dbName), and the calls the Connection() method on the database handle to obtain a handle to the specified collection ("colName")
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance is ready")
}

// MONGODB helpers - file

func insertOneTicket(ticket model.Ticket) {
	// Insert a new ticket into MongoDB
	inserted, err := collection.InsertOne(context.Background(), ticket)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted 1 ticket into the database with ID:", inserted.InsertedID)
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	// This line sets the Content-Type header of the HTTP response to indicate that the response body will contain JSON data.
	w.Header().Set("Content-Type", "application/json")

	// This line declares a variable named 'ticket' of type model.Ticket, which is a struct that likely represents a ticket object
	var ticket model.Ticket
	// This line decodes the JSON request body from the HTTP request('r.Body) into the 'ticket' variable. It uses 'json.NewDecoder()' to create a new JSON decoder and then calls 'Decode()' to decode the JSON data into the 'ticket' variable. Any decoding errors are stores in the 'err' varible
 	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	// This line sets the CreateAt field of the 'ticket' strucit to the current time 
	ticket.CreatedAt = time.Now()
	// This line sets the UpdateAt field of the 'ticket' struct to the current time 
	ticket.UpdatedAt = time.Now()

	// Call insertOneTicket function to insert the ticket into MongoDB
	// It passes the 'ticket' object as an argumewnt to the 'insertOneTicket'
	insertOneTicket(ticket)

	// This line encodes the 'ticket' object into JSON format and writes it as the reponse body of the HTTP response writer ('w') 
	json.NewEncoder(w).Encode(ticket)
}

// updateOneTicket updates a ticket in MongoDB.
func updateOneTicket(id primitive.ObjectID, updatedTicket model.Ticket) error {
	// This line creates a filter to find the ticket by its ID. 
	// It specifies the '_id' field in the MongoDB document
	filter := bson.M{"_id": id}

	// This lione creates an update operation using the '$set' operator, which updates the specified fields of the document without replacing the entire document. It sets the fields  'title', 'description', 'status','assignedTo', 'department', and 'updatedAt' with the corresponding values from the 'updatedTicket' parameter
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTicket.Title,
			"description": updatedTicket.Description,
			"status":      updatedTicket.Status,
			"assignedTo":  updatedTicket.AssignedTo,
			"department":  updatedTicket.Department,
			"updatedAt":   time.Now(),
		},
	}

	// This line performs the updpate operation in MongoDB using the 'UpdateOne' method of the collection. It takes a context, filter, and update as parameters. It returns a result and an error. The result contains information about the update operation, such as the number of documents matched and modified
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Failed to update ticket:", err)
		return err
	}
	// This line logs a success message if the update operation is successful
	log.Println("Ticket updated successfully.")
	return nil
}

// UpdateTicket handles the HTTP PUT request to update a ticket.
func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// This line extracts the URL parameters from the HTTP request using the 'mux.Vars' functuon, which returns a map of route variables for the current request
	vars := mux.Vars(r)

	// This line extracts the ticket ID from the URL parameters and converts it to a 'primitive.ObjectID'
	// If the extraction or conversion fails, its returns a 400 Bad Request error
	ticketID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// Decode the updated ticket data from the request body
	var updatedTicket model.Ticket
	//This line decodes the JSON request body into a 'model.Ticket' struct representing the updated ticket.
	err = json.NewDecoder(r.Body).Decode(&updatedTicket)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// This line calls the 'updateOneTicket' funtion to update the ticket in MongoDB
	//If an error occurs during the update operation, it returns a 500 Internal Server Error
	err = updateOneTicket(ticketID, updatedTicket)
	if err != nil {
		http.Error(w, "Failed to update ticket", http.StatusInternalServerError)
		return
	}

	// This line endocdes the success message as JSON and writes it to the HTTP response body , indicating that the ticket was updated successfully
	json.NewEncoder(w).Encode("Ticket updated successfully")
}

// assignTicketToEmployee assigns a ticket to an employee in MongoDB.
func assignTicketToEmployee(ticketID, employeeID primitive.ObjectID) error {
	// This line creates a fotter to find the ticket by its ID. It specifes the '_id' field in the MongoDB documemt
	filter := bson.M{"_id": ticketID}

	// This line creates an update operation using the '$set' operator, which updates the specified fields of the documennt without replacing the entire document It sets the 'assignedTo' field to the 'employeeID' parameter and updates the 'updatedAt' field with the current time.
	update := bson.M{
		"$set": bson.M{
			"assignedTo": employeeID,
			"updatedAt":  time.Now(),
		},
	}

	// This line performs the update operation in MongoDB using the 'UpdateOne' method of the collection.
	//It takes a context, filter , and update as parameters.
	// It returns a result and an error 
	// The result contains informatuon about the update operation, such as the number of documents matched and modified.
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Failed to assign ticket to employee:", err)
		return err
	}

	// This line logs a success message if the assignment operaton is successful
	log.Println("Ticket assigned to employee successfully.")
	return nil
}

// AssignTicket handles the HTTP PUT request to assign a ticket to an employee.
func AssignTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// This line extracts the URL parameter from the HTTP request using the 'mux.Vars' functuon , which return a map of rute variables for the current request.
	vars := mux.Vars(r)

	// This line extracts the ticket ID from the URL parameters and converts it to a 'primitive.ObjectID'
	// If the extraction or conversion fails, it returns a 400 Bad Request error
	ticketID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	//	This line extracts the employee ID from the URL parameter and con verts it to a 'primitive.ObjectID'
	// If the extraction or conversion fails, it return a 400 Bad Request error
	employeeID, err := primitive.ObjectIDFromHex(vars["employeeID"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	//This line calls the 'assignTicketToEmployee' function to assign the ticket to the employee in MOngoDB 
	// If an error occurs during the assignnment operation, it return a 400 Internal Server Error.
	err = assignTicketToEmployee(ticketID, employeeID)
	if err != nil {
		http.Error(w, "Failed to assign ticket to employee", http.StatusInternalServerError)
		return
	}

	// This line encodes the success message as JSON and writes it to the HTTP response body, indicationg that the ticket was assigned to the employee successfully
	json.NewEncoder(w).Encode("Ticket assigned to employee successfully")
}

// updateTicketStatus updates the status of a ticket in MongoDB.
func updateTicketStatus(ticketID primitive.ObjectID, newStatus string) error {
	// This line creates a filter to find the ticket by its ID. 
	// It specifies the '_id' field in the MOngoDB document
	filter := bson.M{"_id": ticketID}

	// This line creates an update operation using the '$set' operator, which updates te specified fields of the document without replacing the entire document.
	// It sets the 'status' field to the 'newStatus' parameter and updates the 'updateAt' field with the current time 
	update := bson.M{
		"$set": bson.M{
			"status":    newStatus,
			"updatedAt": time.Now(),
		},
	}

	// This line performs the update operation in MongoDB using the 'UpdateOne' method of the collection. It takes a context, filter , and update as parameter.
	// It returns a retust and an error 
	// The result contains information about the update operation, such as the number of documents matched and modified.
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Failed to update ticket status:", err)
		return err
	}

	// This line logs a success message if the status update operation is successful
	log.Println("Ticket status updated successfully.")
	return nil
}

// UpdateTicketStatus handles the HTTP PUT request to update the status of a ticket.
func UpdateTicketStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// This line extracts the URL parameters from the HTTP request using the 'mux.Vars' functuon , which returns a map of route variables for the current request.
	vars := mux.Vars(r)

	// This line extracts the ticket ID from the URL parameter and converts it to a 'primitive.ObjectID' 
	// If the extraction or conversion fails, it returns a 400 Bad Request error
	ticketID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	//This lone declares a new struct 'newStatus' with a 'Status' field , which is used to decode the new status from the request body
	var newStatus struct {
		Status string `json:"status"`
	}

	// This line calls the updateTicketStatus function to update the status of the ticket in MongoDB
	// f an error occurs during the update operation, it returns a 500 Internal Server Error.
	err = json.NewDecoder(r.Body).Decode(&newStatus)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Perform the status update operation
	err = updateTicketStatus(ticketID, newStatus.Status)
	if err != nil {
		http.Error(w, "Failed to update ticket status", http.StatusInternalServerError)
		return
	}

	// This line encodes the success message as JSON and writes it to the HTTP response body, indicating that the ticket status was updated successfully.
	json.NewEncoder(w).Encode("Ticket status updated successfully")
}

// getTicketsByStatus retrieves tickets from MongoDB based on their status.
func getTicketsByStatus(status string) ([]model.Ticket, error) {
	// his line creates a filter to find tickets with the specified status.
	// It specifies the status field in the MongoDB document.
	filter := bson.M{"status": status}

	//This line defines options for querying, such as sorting. In this case, it creates a default options instance for the find operation.
	options := options.Find()

	// This line performs the find operation in MongoDB using the Find method of the collection. It takes a context, filter, and options as parameters. It returns a cursor and an error. The cursor represents the result set of the query.
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		log.Println("Failed to find tickets by status:", err)
		return nil, err
	}
	//This line defers closing the cursor until the surrounding function (getTicketsByStatus) returns.
	//This ensures that the cursor is closed after all processing is done.
	defer cursor.Close(context.Background())

	// This line initializes a slice to store retrieved tickets.
	var tickets []model.Ticket

	// This line iterates over the cursor and decodes each document into a model.Ticket struct
	// Inside the loop, it decodes each document using cursor.Decode(&ticket) and appends the decoded ticket to the tickets slice.
	for cursor.Next(context.Background()) {
		var ticket model.Ticket
		err := cursor.Decode(&ticket)
		if err != nil {
			log.Println("Failed to decode ticket:", err)
			continue
		}
		tickets = append(tickets, ticket)
	}

	// This line checks for any errors encountered during cursor iteration. 
	//If an error is found, it logs the error and returns it.
	if err := cursor.Err(); err != nil {
		log.Println("Error iterating over cursor:", err)
		return nil, err
	}

	log.Println("Retrieved tickets by status successfully.")
	//This line returns the retrieved tickets and a nil error, indicating success.
	return tickets, nil
}

// GetTicketsByStatus handles the HTTP GET request to retrieve tickets by status.
func GetTicketsByStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// This line extracts the status query parameter from the HTTP request URL.
	status := r.URL.Query().Get("status")

	// This line calls the getTicketsByStatus function to retrieve tickets with the specified status from MongoDB.
	tickets, err := getTicketsByStatus(status)
	//This line checks if an error occurred during the retrieval operation
	//If an error is found, it returns a 500 Internal Server Error.
	if err != nil {
		http.Error(w, "Failed to retrieve tickets by status", http.StatusInternalServerError)
		return
	}

	// This line encodes the retrieved tickets as JSON and writes them to the HTTP response body.
	json.NewEncoder(w).Encode(tickets)
}

// getTicketsByDepartment retrieves tickets from MongoDB based on their department.
func getTicketsByDepartment(department string) ([]model.Ticket, error) {
	// This line creates a filter to find tickets with the specified department.
	// It specifies the department field in the MongoDB document.
	filter := bson.M{"department": department}

	// This line defines options for querying, such as sorting
	//In this case, it creates a default options instance for the find operation.
	options := options.Find()

	// This line performs the find operation in MongoDB using the Find method of the collection. It takes a context, filter, and options as parameters.
	// It returns a cursor and an error. The cursor represents the result set of the query.
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		log.Println("Failed to find tickets by department:", err)
		return nil, err
	}
	// This line defers closing the cursor until the surrounding function (getTicketsByDepartment) returns. 
	// This ensures that the cursor is closed after all processing is done.
	defer cursor.Close(context.Background())

	// This line initializes a slice to store retrieved tickets.
	var tickets []model.Ticket

	// This line iterates over the cursor and decodes each document into a model.Ticket struct
	// Inside the loop, it decodes each document using cursor.Decode(&ticket) and appends the decoded ticket to the tickets slice.
	for cursor.Next(context.Background()) {
		var ticket model.Ticket
		err := cursor.Decode(&ticket)
		if err != nil {
			log.Println("Failed to decode ticket:", err)
			continue
		}
		tickets = append(tickets, ticket)
	}

	// This line checks for any errors encountered during cursor iteration
	//  If an error is found, it logs the error and returns it.
	if err := cursor.Err(); err != nil {
		log.Println("Error iterating over cursor:", err)
		return nil, err
	}

	log.Println("Retrieved tickets by department successfully.")
	//This line returns the retrieved tickets and a nil error, indicating success.
	return tickets, nil
}

// GetTicketsByDepartment handles the HTTP GET request to retrieve tickets by department.
func GetTicketsByDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// This line calls the getTicketsByDepartment function to retrieve tickets with the specified department from MongoDB.
	department := r.URL.Query().Get("department")

	// This line checks if an error occurred during the retrieval operation
	// If an error is found, it returns a 500 Internal Server Error.
	tickets, err := getTicketsByDepartment(department)
	if err != nil {
		http.Error(w, "Failed to retrieve tickets by department", http.StatusInternalServerError)
		return
	}

	// This line encodes the retrieved tickets as JSON and writes them to the HTTP response body.
	json.NewEncoder(w).Encode(tickets)
}

// getTicketByID retrieves a ticket from MongoDB by its ID.
func getTicketByID(id primitive.ObjectID) (*model.Ticket, error) {
	// This line creates a filter to find a ticket by its ID.
	// It specifies the _id field in the MongoDB document.
	filter := bson.M{"_id": id}

	// Perform the find operation in MongoDB
	var ticket model.Ticket
	// This line performs the find operation in MongoDB using the FindOne method of the collection.
	// It takes a context and a filter as parameters.
	// It retrieves a single document that matches the filter and decodes it into the ticket variable.

	err := collection.FindOne(context.Background(), filter).Decode(&ticket)
	//This line checks if an error occurred during the find operation.
	// f an error is found, it logs the error and returns it.
	if err != nil {
		log.Println("Failed to find ticket by ID:", err)
		return nil, err
	}

	log.Println("Retrieved ticket details successfully.")
	//This line returns a pointer to the retrieved ticket and a nil error, indicating success.
	return &ticket, nil
}

// GetTicketDetails handles the HTTP GET request to retrieve details of a ticket.
func GetTicketDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// This line extracts the URL parameters from the HTTP request using the mux.Vars function, which returns a map of route variables for the current request.
	vars := mux.Vars(r)
	//This line extracts the ticket ID from the URL parameters and converts it to a primitive.ObjectID
	//If the extraction or conversion fails, it returns a 400 Bad Request error.
	ticketID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// This line calls the getTicketByID function to retrieve details of the ticket with the specified ID from MongoDB.
	ticket, err := getTicketByID(ticketID)
	// This line checks if an error occurred during the retrieval operation. If an error is found, it returns a 500 Internal Server Error.
	if err != nil {
		http.Error(w, "Failed to retrieve ticket details", http.StatusInternalServerError)
		return
	}

	// This line encodes the retrieved ticket details as JSON and writes them to the HTTP response body.
	json.NewEncoder(w).Encode(ticket)
} 

// deleteTicket deletes a ticket from MongoDB by its ID.
func deleteTicket(id primitive.ObjectID) error {
	// This line creates a filter to find the ticket by its ID.
	// It specifies the _id field in the MongoDB document.
	filter := bson.M{"_id": id}

	// This line performs the delete operation in MongoDB using the DeleteOne method of the collection.
	// It takes a context and a filter as parameters. 
	// It deletes a single document that matches the filter.
	_, err := collection.DeleteOne(context.Background(), filter)
	//This line checks if an error occurred during the delete operation.
	// If an error is found, it logs the error and returns it.
	if err != nil {
		log.Println("Failed to delete ticket:", err)
		return err
	}

	log.Println("Ticket deleted successfully.")
	//This line returns nil, indicating success.
	return nil
}

// DeleteTicket handles the HTTP DELETE request to delete a ticket.
func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// This line extracts the URL parameters from the HTTP request using the mux.Vars function, which returns a map of route variables for the current request.
	vars := mux.Vars(r)
	//This line extracts the ticket ID from the URL parameters and converts it to a primitive.ObjectID
	// If the extraction or conversion fails, it returns a 400 Bad Request error.
	ticketID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// This line calls the deleteTicket function to delete the ticket with the specified ID from MongoDB.
	err = deleteTicket(ticketID)
	// This line checks if an error occurred during the deletion operation.
	// If an error is found, it returns a 500 Internal Server Error.
	if err != nil {
		http.Error(w, "Failed to delete ticket", http.StatusInternalServerError)
		return
	}

	// This line encodes the success message as JSON and writes it to the HTTP response body, indicating that the ticket was deleted successfully.
	json.NewEncoder(w).Encode("Ticket deleted successfully")
}

// createSubtickets creates sub-tickets for a given parent ticket in MongoDB.
func createSubtickets(parentTicketID primitive.ObjectID, subTickets []model.Ticket) error {
	// Set creation and update time for each sub-ticket
	currentTime := time.Now()
	for i := range subTickets {
		subTickets[i].CreatedAt = currentTime
		subTickets[i].UpdatedAt = currentTime
	}

	// Convert the slice of model.Ticket to slice of interface{}
	var interfaceSlice []interface{}
	for _, ticket := range subTickets {
   	 interfaceSlice = append(interfaceSlice, ticket)
	}

// Perform the insert operation in MongoDB
_, err := collection.InsertMany(context.Background(), interfaceSlice)
	if err != nil {
		log.Println("Failed to create sub-tickets:", err)
		return err
	}

	log.Println("Sub-tickets created successfully.")
	return nil
}

// CreateSubtickets handles the HTTP POST request to create sub-tickets.
func CreateSubtickets(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    // Decode the sub-tickets data from the request body
    var subTickets []model.Ticket
    err := json.NewDecoder(r.Body).Decode(&subTickets)
    if err != nil {
        http.Error(w, "Failed to decode request body", http.StatusBadRequest)
        return
    }

    // Convert the slice of model.Ticket to slice of interface{}
    var interfaceSlice []interface{}
    for _, ticket := range subTickets {
        interfaceSlice = append(interfaceSlice, ticket)
    }

    // Perform the insert operation in MongoDB
    _, err = collection.InsertMany(context.Background(), interfaceSlice)
    if err != nil {
        http.Error(w, "Failed to create sub-tickets", http.StatusInternalServerError)
        return
    }

    // Return success response
    json.NewEncoder(w).Encode("Sub-tickets created successfully")
}
