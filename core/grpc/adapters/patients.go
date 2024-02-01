package adapters

import (
	"context"
	pb "gocourse17/core/grpc/generated"
	"gocourse17/core/service/patients"
)

func NewPatients(s *patients.Service) *Patients {
	return &Patients{service: s}
}

type Patients struct {
	service *patients.Service
	pb.UnimplementedPatientServiceServer
}

func (s *Patients) AddPatient(ctx context.Context, in *pb.AddPatientRequest) (*pb.AddPatientResponse, error) {
	patientEntity := patients.Patient{
		in.GetPatient().Id,
		in.GetPatient().Name,
		in.GetPatient().Age,
		in.GetPatient().Diagnosis,
	}

	if _, err := s.service.AddPatient(ctx, &patientEntity); err != nil {
		return nil, err
	} else {
		return &pb.AddPatientResponse{Message: "Patient added successfully"}, nil
	}
}

func (s *Patients) GetPatient(ctx context.Context, in *pb.GetPatientRequest) (*pb.GetPatientResponse, error) {
	if res, err := s.service.GetPatient(ctx, in.GetId()); err != nil {
		return nil, err
	} else {
		return &pb.GetPatientResponse{Patient: &pb.Patient{
			Id:        res.ID,
			Name:      res.Name,
			Age:       res.Age,
			Diagnosis: res.Diagnosis,
		}}, nil
	}
}

func (s *Patients) UpdatePatient(ctx context.Context, in *pb.UpdatePatientRequest) (*pb.UpdatePatientResponse, error) {
	patientEntity := patients.Patient{
		in.GetPatient().Id,
		in.GetPatient().Name,
		in.GetPatient().Age,
		in.GetPatient().Diagnosis,
	}
	if _, err := s.service.UpdatePatient(ctx, in.GetPatient().GetId(), &patientEntity); err != nil {
		return nil, err
	} else {

		return &pb.UpdatePatientResponse{Message: "Patient updated successfully"}, nil
	}
}
