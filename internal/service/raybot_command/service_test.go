package raybotcommand_test

import (
	"context"
	"io"
	"log/slog"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/tuanvumaihuynh/roboflow/internal/model"
	"github.com/tuanvumaihuynh/roboflow/internal/model/mocks"
	raybotcommand "github.com/tuanvumaihuynh/roboflow/internal/service/raybot_command"
	"github.com/tuanvumaihuynh/roboflow/pkg/paging"
	"github.com/tuanvumaihuynh/roboflow/pkg/xsort"
)

func TestRaybotCommandService(t *testing.T) {
	log := slog.New(slog.NewTextHandler(io.Discard, nil))
	ctx := context.Background()

	// t.Run("Create", func(t *testing.T) {
	// 	type testCase struct {
	// 		name              string
	// 		cmd               raybotcommand.CreateRaybotCommandCommand
	// 		raybotCommandRepo *mocks.FakeRaybotCommandRepository
	// 		raybotRepo        *mocks.FakeRaybotRepository
	// 		qrLocationRepo    *mocks.FakeQRLocationRepository
	// 		eventPublisher    *pubsubmocks.FakePublisher
	// 		mockBehavior      func(*testCase)
	// 		shouldErr         bool
	// 	}

	// 	testCases := []testCase{
	// 		{
	// 			name: "Should create successfully",
	// 			cmd: raybotcommand.CreateRaybotCommandCommand{
	// 				RaybotID: validID,
	// 				Type:     model.RaybotCommandTypeMoveForward,
	// 				Input:    map[string]interface{}{},
	// 			},
	// 			raybotCommandRepo: mocks.NewFakeRaybotCommandRepository(t),
	// 			raybotRepo:        mocks.NewFakeRaybotRepository(t),
	// 			qrLocationRepo:    mocks.NewFakeQRLocationRepository(t),
	// 			eventPublisher:    pubsubmocks.NewFakePublisher(t),
	// 			mockBehavior: func(tc *testCase) {
	// 				tc.raybotRepo.On("GetState", ctx, validID).Return(model.RaybotStatusIdle, nil)
	// 				tc.raybotCommandRepo.On("Create", ctx, mock.Anything).Return(nil)
	// 				tc.eventPublisher.On("Publish", "raybot.command.created", mock.Anything).Return(nil)
	// 			},
	// 			shouldErr: false,
	// 		},
	// 		{
	// 			name: "Should validate command before do anything",
	// 			cmd: raybotcommand.CreateRaybotCommandCommand{
	// 				RaybotID: uuid.Nil,
	// 				Type:     model.RaybotCommandTypeMoveForward,
	// 				Input:    map[string]interface{}{},
	// 			},
	// 			raybotCommandRepo: mocks.NewFakeRaybotCommandRepository(t),
	// 			raybotRepo:        mocks.NewFakeRaybotRepository(t),
	// 			qrLocationRepo:    mocks.NewFakeQRLocationRepository(t),
	// 			eventPublisher:    pubsubmocks.NewFakePublisher(t),
	// 			mockBehavior: func(tc *testCase) {
	// 				tc.raybotRepo.AssertNotCalled(t, "GetState", ctx, uuid.Nil)
	// 				tc.raybotCommandRepo.AssertNotCalled(t, "Create", ctx, mock.Anything)
	// 				tc.eventPublisher.AssertNotCalled(t, "Publish", "raybot.command.created", mock.Anything)
	// 			},
	// 			shouldErr: true,
	// 		},
	// 		{
	// 			name: "Should return error when get state fails",
	// 			cmd: raybotcommand.CreateRaybotCommandCommand{
	// 				RaybotID: validID,
	// 				Type:     model.RaybotCommandTypeMoveForward,
	// 				Input:    map[string]interface{}{},
	// 			},
	// 			raybotCommandRepo: mocks.NewFakeRaybotCommandRepository(t),
	// 			raybotRepo:        mocks.NewFakeRaybotRepository(t),
	// 			qrLocationRepo:    mocks.NewFakeQRLocationRepository(t),
	// 			eventPublisher:    pubsubmocks.NewFakePublisher(t),
	// 			mockBehavior: func(tc *testCase) {
	// 				tc.raybotRepo.On("GetState", ctx, validID).Return(model.RaybotStatus(""), assert.AnError)
	// 				tc.raybotCommandRepo.AssertNotCalled(t, "Create", ctx, mock.Anything)
	// 				tc.eventPublisher.AssertNotCalled(t, "Publish", "raybot.command.created", mock.Anything)
	// 			},
	// 			shouldErr: true,
	// 		},
	// 		{
	// 			name: "Should return error when raybot is OFFLINE",
	// 			cmd: raybotcommand.CreateRaybotCommandCommand{
	// 				RaybotID: validID,
	// 				Type:     model.RaybotCommandTypeMoveForward,
	// 				Input:    map[string]interface{}{},
	// 			},
	// 			raybotCommandRepo: mocks.NewFakeRaybotCommandRepository(t),
	// 			raybotRepo:        mocks.NewFakeRaybotRepository(t),
	// 			qrLocationRepo:    mocks.NewFakeQRLocationRepository(t),
	// 			eventPublisher:    pubsubmocks.NewFakePublisher(t),
	// 			mockBehavior: func(tc *testCase) {
	// 				tc.raybotRepo.On("GetState", ctx, validID).Return(model.RaybotStatusOffline, nil)
	// 				tc.raybotCommandRepo.AssertNotCalled(t, "Create", ctx, mock.Anything)
	// 				tc.eventPublisher.AssertNotCalled(t, "Publish", "raybot.command.created", mock.Anything)
	// 			},
	// 			shouldErr: true,
	// 		},
	// 		{
	// 			name: "Should return error when raybot is BUSY and command is not STOP",
	// 			cmd: raybotcommand.CreateRaybotCommandCommand{
	// 				RaybotID: validID,
	// 				Type:     model.RaybotCommandTypeMoveForward,
	// 				Input:    map[string]interface{}{},
	// 			},
	// 			raybotCommandRepo: mocks.NewFakeRaybotCommandRepository(t),
	// 			raybotRepo:        mocks.NewFakeRaybotRepository(t),
	// 			qrLocationRepo:    mocks.NewFakeQRLocationRepository(t),
	// 			eventPublisher:    pubsubmocks.NewFakePublisher(t),
	// 			mockBehavior: func(tc *testCase) {
	// 				tc.raybotRepo.On("GetState", ctx, validID).Return(model.RaybotStatusBusy, nil)
	// 				tc.raybotCommandRepo.AssertNotCalled(t, "Create", ctx, mock.Anything)
	// 				tc.eventPublisher.AssertNotCalled(t, "Publish", "raybot.command.created", mock.Anything)
	// 			},
	// 			shouldErr: true,
	// 		},
	// 		{
	// 			name: "Should return error when repository fails to create",
	// 			cmd: raybotcommand.CreateRaybotCommandCommand{
	// 				RaybotID: validID,
	// 				Type:     model.RaybotCommandTypeMoveForward,
	// 				Input:    map[string]interface{}{},
	// 			},
	// 			raybotCommandRepo: mocks.NewFakeRaybotCommandRepository(t),
	// 			raybotRepo:        mocks.NewFakeRaybotRepository(t),
	// 			qrLocationRepo:    mocks.NewFakeQRLocationRepository(t),
	// 			eventPublisher:    pubsubmocks.NewFakePublisher(t),
	// 			mockBehavior: func(tc *testCase) {
	// 				tc.raybotRepo.On("GetState", ctx, validID).Return(model.RaybotStatusIdle, nil)
	// 				tc.raybotCommandRepo.On("Create", ctx, mock.Anything).Return(assert.AnError)
	// 				tc.eventPublisher.AssertNotCalled(t, "Publish", "raybot.command.created", mock.Anything)
	// 			},
	// 			shouldErr: true,
	// 		},
	// 		{
	// 			name: "Should return error when event publisher fails to publish",
	// 			cmd: raybotcommand.CreateRaybotCommandCommand{
	// 				RaybotID: validID,
	// 				Type:     model.RaybotCommandTypeMoveForward,
	// 				Input:    map[string]interface{}{},
	// 			},
	// 			raybotCommandRepo: mocks.NewFakeRaybotCommandRepository(t),
	// 			raybotRepo:        mocks.NewFakeRaybotRepository(t),
	// 			qrLocationRepo:    mocks.NewFakeQRLocationRepository(t),
	// 			eventPublisher:    pubsubmocks.NewFakePublisher(t),
	// 			mockBehavior: func(tc *testCase) {
	// 				tc.raybotRepo.On("GetState", ctx, validID).Return(model.RaybotStatusIdle, nil)
	// 				tc.raybotCommandRepo.On("Create", ctx, mock.Anything).Return(nil)
	// 				tc.eventPublisher.On("Publish", "raybot.command.created", mock.Anything).Return(assert.AnError)
	// 			},
	// 			shouldErr: true,
	// 		},
	// 	}

	// 	for _, tc := range testCases {
	// 		t.Run(tc.name, func(t *testing.T) {
	// 			tc.mockBehavior(&tc)

	// 			s := raybotcommand.NewService(tc.raybotCommandRepo, tc.raybotRepo, tc.qrLocationRepo, tc.eventPublisher, log)
	// 			result, err := s.Create(ctx, tc.cmd)

	// 			if tc.shouldErr {
	// 				assert.Error(t, err)
	// 			} else {
	// 				assert.Equal(t, tc.cmd.RaybotID, result.RaybotID)
	// 				assert.Equal(t, tc.cmd.Type, result.Type)
	// 				assert.Equal(t, tc.cmd.Input, result.Input)
	// 				assert.NoError(t, err)
	// 			}
	// 		})
	// 	}
	// })
	t.Run("Delete", func(t *testing.T) {
		testCases := []struct {
			name         string
			cmd          raybotcommand.DeleteRaybotCommandCommand
			mockBehavior func(*mocks.FakeRaybotCommandRepository)
			shouldErr    bool
		}{
			{
				name: "Should delete successfully",
				cmd: raybotcommand.DeleteRaybotCommandCommand{
					ID: validID,
				},
				mockBehavior: func(r *mocks.FakeRaybotCommandRepository) {
					r.On("Delete", ctx, validID).Return(nil)
				},
				shouldErr: false,
			}, {
				name: "Should validate command before do anything",
				cmd: raybotcommand.DeleteRaybotCommandCommand{
					ID: uuid.Nil,
				},
				mockBehavior: func(r *mocks.FakeRaybotCommandRepository) {
					r.AssertNotCalled(t, "Delete", ctx, uuid.Nil)
				},
				shouldErr: true,
			}, {
				name: "Should return error when repository fails to delete",
				cmd: raybotcommand.DeleteRaybotCommandCommand{
					ID: validID,
				},
				mockBehavior: func(r *mocks.FakeRaybotCommandRepository) {
					r.On("Delete", ctx, validID).Return(assert.AnError)
				},
				shouldErr: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				repo := &mocks.FakeRaybotCommandRepository{}
				tc.mockBehavior(repo)

				s := raybotcommand.NewService(repo, nil, nil, nil, log)
				err := s.Delete(ctx, tc.cmd)

				if tc.shouldErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})
	t.Run("GetByID", func(t *testing.T) {
		testCases := []struct {
			name         string
			q            raybotcommand.GetRaybotCommandByIDQuery
			mockBehavior func(*mocks.FakeRaybotCommandRepository)
			shouldErr    bool
		}{
			{
				name: "Should get successfully",
				q: raybotcommand.GetRaybotCommandByIDQuery{
					ID: validID,
				},
				mockBehavior: func(r *mocks.FakeRaybotCommandRepository) {
					r.On("Get", ctx, validID).Return(validRaybotCommand, nil)
				},
				shouldErr: false,
			}, {
				name: "Should validate query before do anything",
				q: raybotcommand.GetRaybotCommandByIDQuery{
					ID: uuid.Nil,
				},
				mockBehavior: func(r *mocks.FakeRaybotCommandRepository) {
					r.AssertNotCalled(t, "Get", ctx, uuid.Nil)
				},
				shouldErr: true,
			}, {
				name: "Should return error when repository fails to get",
				q: raybotcommand.GetRaybotCommandByIDQuery{
					ID: validID,
				},
				mockBehavior: func(r *mocks.FakeRaybotCommandRepository) {
					r.On("Get", ctx, validID).Return(validRaybotCommand, assert.AnError)
				},
				shouldErr: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				raybotCommandRepo := mocks.NewFakeRaybotCommandRepository(t)
				tc.mockBehavior(raybotCommandRepo)

				s := raybotcommand.NewService(raybotCommandRepo, nil, nil, nil, log)
				_, err := s.GetByID(ctx, tc.q)

				if tc.shouldErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})
	t.Run("List", func(t *testing.T) {
		validPagingListRaybotCommands := paging.List[model.RaybotCommand]{
			Items:     []model.RaybotCommand{validRaybotCommand},
			TotalItem: 1,
		}

		testCases := []struct {
			name         string
			q            raybotcommand.ListRaybotCommandQuery
			mockBehavior func(*mocks.FakeRaybotCommandRepository)
			shouldErr    bool
		}{
			{
				name: "Should list successfully",
				q: raybotcommand.ListRaybotCommandQuery{
					RaybotID:     validID,
					PagingParams: paging.NewParams(nil, nil),
					Sorts:        []xsort.Sort{},
				},
				mockBehavior: func(r *mocks.FakeRaybotCommandRepository) {
					r.On("List", ctx, mock.Anything, mock.Anything, mock.Anything).Return(&validPagingListRaybotCommands, nil)
				},
				shouldErr: false,
			},
			{
				name: "Should validate query before do anything",
				q: raybotcommand.ListRaybotCommandQuery{
					RaybotID:     uuid.Nil,
					PagingParams: paging.NewParams(nil, nil),
					Sorts:        []xsort.Sort{},
				},
				mockBehavior: func(r *mocks.FakeRaybotCommandRepository) {
					r.AssertNotCalled(t, "List", ctx, uuid.Nil, mock.Anything, mock.Anything)
				},
				shouldErr: true,
			},
			{
				name: "Should return error when repository fails to list",
				q: raybotcommand.ListRaybotCommandQuery{
					RaybotID:     validID,
					PagingParams: paging.NewParams(nil, nil),
					Sorts:        []xsort.Sort{},
				},
				mockBehavior: func(r *mocks.FakeRaybotCommandRepository) {
					r.On("List", ctx, validID, mock.Anything, mock.Anything).Return(&validPagingListRaybotCommands, assert.AnError)
				},
				shouldErr: true,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				raybotCommandRepo := mocks.NewFakeRaybotCommandRepository(t)
				tc.mockBehavior(raybotCommandRepo)

				s := raybotcommand.NewService(raybotCommandRepo, nil, nil, nil, log)
				_, err := s.List(ctx, tc.q)

				if tc.shouldErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			})
		}
	})
}

var (
	validID            = uuid.New()
	validRaybotCommand = model.RaybotCommand{
		RaybotID:    validID,
		ID:          validID,
		Type:        model.RaybotCommandTypeCheckQrCode,
		Status:      model.RaybotCommandStatusPending,
		Input:       map[string]interface{}{},
		Output:      map[string]interface{}{},
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
)