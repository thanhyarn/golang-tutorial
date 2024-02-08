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

const connectionString = "mongodb+srv://thanhyarn:123@atlascluster.bfqeyvn.mongodb.net/TestGoLang?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

//MOST IMPORTANT
var collection *mongo.Collection

// connect with monogoDB

func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(colName)

	//collection instance
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
	w.Header().Set("Content-Type", "application/json")
	var ticket model.Ticket
	// Decode request body into a Ticket struct
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Set creation and update time for the ticket
	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()

	// Call insertOneTicket function to insert the ticket into MongoDB
	insertOneTicket(ticket)

	// Return information of the created ticket
	json.NewEncoder(w).Encode(ticket)
}

// updateOneTicket updates a ticket in MongoDB.
func updateOneTicket(id primitive.ObjectID, updatedTicket model.Ticket) error {
	// Create filter to find the ticket by its ID
	filter := bson.M{"_id": id}

	// Create update to set the new ticket data
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

	// Perform the update operation in MongoDB
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Failed to update ticket:", err)
		return err
	}

	log.Println("Ticket updated successfully.")
	return nil
}

// UpdateTicket handles the HTTP PUT request to update a ticket.
func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the ticket ID from the request URL
	vars := mux.Vars(r)
	ticketID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// Decode the updated ticket data from the request body
	var updatedTicket model.Ticket
	err = json.NewDecoder(r.Body).Decode(&updatedTicket)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Perform the update operation
	err = updateOneTicket(ticketID, updatedTicket)
	if err != nil {
		http.Error(w, "Failed to update ticket", http.StatusInternalServerError)
		return
	}

	// Return success response
	json.NewEncoder(w).Encode("Ticket updated successfully")
}

// assignTicketToEmployee assigns a ticket to an employee in MongoDB.
func assignTicketToEmployee(ticketID, employeeID primitive.ObjectID) error {
	// Create filter to find the ticket by its ID
	filter := bson.M{"_id": ticketID}

	// Create update to set the assignedTo field to the employee's ID
	update := bson.M{
		"$set": bson.M{
			"assignedTo": employeeID,
			"updatedAt":  time.Now(),
		},
	}

	// Perform the update operation in MongoDB
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Failed to assign ticket to employee:", err)
		return err
	}

	log.Println("Ticket assigned to employee successfully.")
	return nil
}

// AssignTicket handles the HTTP PUT request to assign a ticket to an employee.
func AssignTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the ticket ID and employee ID from the request URL
	vars := mux.Vars(r)
	ticketID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}
	employeeID, err := primitive.ObjectIDFromHex(vars["employeeID"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Perform the assignment operation
	err = assignTicketToEmployee(ticketID, employeeID)
	if err != nil {
		http.Error(w, "Failed to assign ticket to employee", http.StatusInternalServerError)
		return
	}

	// Return success response
	json.NewEncoder(w).Encode("Ticket assigned to employee successfully")
}

// updateTicketStatus updates the status of a ticket in MongoDB.
func updateTicketStatus(ticketID primitive.ObjectID, newStatus string) error {
	// Create filter to find the ticket by its ID
	filter := bson.M{"_id": ticketID}

	// Create update to set the new status
	update := bson.M{
		"$set": bson.M{
			"status":    newStatus,
			"updatedAt": time.Now(),
		},
	}

	// Perform the update operation in MongoDB
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println("Failed to update ticket status:", err)
		return err
	}

	log.Println("Ticket status updated successfully.")
	return nil
}

// UpdateTicketStatus handles the HTTP PUT request to update the status of a ticket.
func UpdateTicketStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the ticket ID from the request URL
	vars := mux.Vars(r)
	ticketID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// Decode the new status from the request body
	var newStatus struct {
		Status string `json:"status"`
	}
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

	// Return success response
	json.NewEncoder(w).Encode("Ticket status updated successfully")
}

// getTicketsByStatus retrieves tickets from MongoDB based on their status.
func getTicketsByStatus(status string) ([]model.Ticket, error) {
	// Create filter to find tickets with the specified status
	filter := bson.M{"status": status}

	// Define options for querying (e.g., sorting)
	options := options.Find()

	// Perform the find operation in MongoDB
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		log.Println("Failed to find tickets by status:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Initialize slice to store retrieved tickets
	var tickets []model.Ticket

	// Iterate over the cursor and decode each document into a Ticket struct
	for cursor.Next(context.Background()) {
		var ticket model.Ticket
		err := cursor.Decode(&ticket)
		if err != nil {
			log.Println("Failed to decode ticket:", err)
			continue
		}
		tickets = append(tickets, ticket)
	}

	// Check for any errors encountered during cursor iteration
	if err := cursor.Err(); err != nil {
		log.Println("Error iterating over cursor:", err)
		return nil, err
	}

	log.Println("Retrieved tickets by status successfully.")
	return tickets, nil
}

// GetTicketsByStatus handles the HTTP GET request to retrieve tickets by status.
func GetTicketsByStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the status query parameter from the request URL
	status := r.URL.Query().Get("status")

	// Perform the query operation
	tickets, err := getTicketsByStatus(status)
	if err != nil {
		http.Error(w, "Failed to retrieve tickets by status", http.StatusInternalServerError)
		return
	}

	// Return retrieved tickets as JSON response
	json.NewEncoder(w).Encode(tickets)
}

// getTicketsByDepartment retrieves tickets from MongoDB based on their department.
func getTicketsByDepartment(department string) ([]model.Ticket, error) {
	// Create filter to find tickets with the specified department
	filter := bson.M{"department": department}

	// Define options for querying (e.g., sorting)
	options := options.Find()

	// Perform the find operation in MongoDB
	cursor, err := collection.Find(context.Background(), filter, options)
	if err != nil {
		log.Println("Failed to find tickets by department:", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Initialize slice to store retrieved tickets
	var tickets []model.Ticket

	// Iterate over the cursor and decode each document into a Ticket struct
	for cursor.Next(context.Background()) {
		var ticket model.Ticket
		err := cursor.Decode(&ticket)
		if err != nil {
			log.Println("Failed to decode ticket:", err)
			continue
		}
		tickets = append(tickets, ticket)
	}

	// Check for any errors encountered during cursor iteration
	if err := cursor.Err(); err != nil {
		log.Println("Error iterating over cursor:", err)
		return nil, err
	}

	log.Println("Retrieved tickets by department successfully.")
	return tickets, nil
}

// GetTicketsByDepartment handles the HTTP GET request to retrieve tickets by department.
func GetTicketsByDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the department query parameter from the request URL
	department := r.URL.Query().Get("department")

	// Perform the query operation
	tickets, err := getTicketsByDepartment(department)
	if err != nil {
		http.Error(w, "Failed to retrieve tickets by department", http.StatusInternalServerError)
		return
	}

	// Return retrieved tickets as JSON response
	json.NewEncoder(w).Encode(tickets)
}

// getTicketByID retrieves a ticket from MongoDB by its ID.
func getTicketByID(id primitive.ObjectID) (*model.Ticket, error) {
	// Create filter to find the ticket by its ID
	filter := bson.M{"_id": id}

	// Perform the find operation in MongoDB
	var ticket model.Ticket
	err := collection.FindOne(context.Background(), filter).Decode(&ticket)
	if err != nil {
		log.Println("Failed to find ticket by ID:", err)
		return nil, err
	}

	log.Println("Retrieved ticket details successfully.")
	return &ticket, nil
}

// GetTicketDetails handles the HTTP GET request to retrieve details of a ticket.
func GetTicketDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the ticket ID from the request URL
	vars := mux.Vars(r)
	ticketID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// Perform the query operation
	ticket, err := getTicketByID(ticketID)
	if err != nil {
		http.Error(w, "Failed to retrieve ticket details", http.StatusInternalServerError)
		return
	}

	// Return retrieved ticket details as JSON response
	json.NewEncoder(w).Encode(ticket)
} 

// deleteTicket deletes a ticket from MongoDB by its ID.
func deleteTicket(id primitive.ObjectID) error {
	// Create filter to find the ticket by its ID
	filter := bson.M{"_id": id}

	// Perform the delete operation in MongoDB
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println("Failed to delete ticket:", err)
		return err
	}

	log.Println("Ticket deleted successfully.")
	return nil
}

// DeleteTicket handles the HTTP DELETE request to delete a ticket.
func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract the ticket ID from the request URL
	vars := mux.Vars(r)
	ticketID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	// Perform the delete operation
	err = deleteTicket(ticketID)
	if err != nil {
		http.Error(w, "Failed to delete ticket", http.StatusInternalServerError)
		return
	}

	// Return success response
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
