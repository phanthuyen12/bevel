package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Course struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Instructor  string `json:"instructor"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (s *SmartContract) Init(ctx contractapi.TransactionContextInterface) error {
	courses := []Course{
		{
			ID:          "course-001",
			Title:       "Hyperledger Fabric Basics",
			Instructor:  "admin",
			Description: "Starter course for Smart LMS Fabric network",
			Status:      "active",
		},
	}

	for _, course := range courses {
		data, err := json.Marshal(course)
		if err != nil {
			return err
		}

		if err := ctx.GetStub().PutState(course.ID, data); err != nil {
			return fmt.Errorf("failed to put course %s: %w", course.ID, err)
		}
	}

	return nil
}

func (s *SmartContract) CreateCourse(
	ctx contractapi.TransactionContextInterface,
	id string,
	title string,
	instructor string,
	description string,
) error {
	exists, err := s.CourseExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("course %s already exists", id)
	}

	course := Course{
		ID:          id,
		Title:       title,
		Instructor:  instructor,
		Description: description,
		Status:      "active",
	}

	data, err := json.Marshal(course)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, data)
}

func (s *SmartContract) ReadCourse(ctx contractapi.TransactionContextInterface, id string) (*Course, error) {
	data, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read course %s: %w", id, err)
	}
	if data == nil {
		return nil, fmt.Errorf("course %s does not exist", id)
	}

	var course Course
	if err := json.Unmarshal(data, &course); err != nil {
		return nil, err
	}

	return &course, nil
}

func (s *SmartContract) UpdateCourse(
	ctx contractapi.TransactionContextInterface,
	id string,
	title string,
	instructor string,
	description string,
	status string,
) error {
	exists, err := s.CourseExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("course %s does not exist", id)
	}

	course := Course{
		ID:          id,
		Title:       title,
		Instructor:  instructor,
		Description: description,
		Status:      status,
	}

	data, err := json.Marshal(course)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, data)
}

func (s *SmartContract) DeleteCourse(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.CourseExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("course %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

func (s *SmartContract) GetAllCourses(ctx contractapi.TransactionContextInterface) ([]*Course, error) {
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var courses []*Course
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var course Course
		if err := json.Unmarshal(queryResponse.Value, &course); err != nil {
			return nil, err
		}

		courses = append(courses, &course)
	}

	return courses, nil
}

func (s *SmartContract) CourseExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	data, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read course %s: %w", id, err)
	}

	return data != nil, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		panic(fmt.Sprintf("failed to create smartlms chaincode: %s", err))
	}

	if err := chaincode.Start(); err != nil {
		panic(fmt.Sprintf("failed to start smartlms chaincode: %s", err))
	}
}
