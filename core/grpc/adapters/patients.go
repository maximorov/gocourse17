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

func (s *Patients) Add(ctx context.Context, in *pb.AddPatientRequest) (*pb.AddPatientResponse, error) {
	patientEntity := patients.Patient{
		in.GetPatient().Id,
		in.GetPatient().Name,
		in.GetPatient().Age,
		in.GetPatient().Diagnosis,
	}

	if _, err := s.service.AddPatient(ctx, &patientEntity); err != nil {
		return &pb.AddPatientResponse{Message: "Patient added successfully"}, nil
	} else {
		return nil, err
	}
}

func (s *Patients) Get(ctx context.Context, in *pb.GetPatientRequest) (*pb.GetPatientResponse, error) {
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

func (s *Patients) Update(ctx context.Context, in *pb.UpdatePatientRequest) (*pb.UpdatePatientResponse, error) {
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
