package seasons

import (
	"context"

	v4Client "github.com/gubarz/gohtb/httpclient/v4"
	"github.com/gubarz/gohtb/internal/common"
	"github.com/gubarz/gohtb/internal/service"
)

func NewService(client service.Client) *Service {
	return &Service{
		base: service.NewBase(client),
	}
}

// Season returns a handle for a specific season with the given ID.
// This handle can be used to perform operations related to that season,
// such as retrieving rewards, user rankings, and follower information.
func (s *Service) Season(id int) *Handle {
	return &Handle{
		client: s.base.Client,
		id:     id,
	}
}

// Rewards retrieves the rewards available for the specified season.
// This includes information about prizes, achievements, and other rewards
// that can be earned during the season.
//
// Example:
//
//	rewards, err := client.Seasons.Season(123).Rewards(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, reward := range rewards.Data {
//		fmt.Printf("Reward: %s (Points: %d)\n", reward.Name, reward.Points)
//	}
func (h *Handle) Rewards(ctx context.Context) (RewardsResponse, error) {
	resp, err := h.client.V4().GetSeasonRewards(h.client.Limiter().Wrap(ctx), h.id)
	if err != nil {
		return RewardsResponse{ResponseMeta: common.ResponseMeta{}}, err
	}

	parsed, meta, err := common.Parse(resp, v4Client.ParseGetSeasonRewardsResponse)
	if err != nil {
		return RewardsResponse{ResponseMeta: meta}, err
	}

	return RewardsResponse{
		Data:         parsed.JSON200.Data,
		ResponseMeta: meta,
	}, nil
}

// UserRank retrieves the current user's ranking information for the specified season.
// This includes position, points, and other ranking details for the authenticated user.
//
// Example:
//
//	rank, err := client.Seasons.Season(123).UserRank(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Current rank: %d (Points: %d)\n", rank.Data.Position, rank.Data.Points)
func (h *Handle) UserRank(ctx context.Context) (UserRankResponse, error) {
	resp, err := h.client.V4().GetSeasonUserRank(h.client.Limiter().Wrap(ctx), h.id)
	if err != nil {
		return UserRankResponse{ResponseMeta: common.ResponseMeta{}}, err
	}

	parsed, meta, err := common.Parse(resp, v4Client.ParseGetSeasonUserRankResponse)
	if err != nil {
		return UserRankResponse{ResponseMeta: meta}, err
	}

	return UserRankResponse{
		Data:         parsed.JSON200.Data,
		ResponseMeta: meta,
	}, nil
}

// UserFollowers retrieves follower information for the current user in the specified season.
// This includes details about users following the authenticated user during the season.
//
// Example:
//
//	followers, err := client.Seasons.Season(123).UserFollowers(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Followers: %d\n", len(followers.Data.Followers))
func (h *Handle) UserFollowers(ctx context.Context) (UserFollowersResponse, error) {
	resp, err := h.client.V4().GetSeasonUserFollowers(h.client.Limiter().Wrap(ctx), h.id)
	if err != nil {
		return UserFollowersResponse{ResponseMeta: common.ResponseMeta{}}, err
	}

	parsed, meta, err := common.Parse(resp, v4Client.ParseGetSeasonUserFollowersResponse)
	if err != nil {
		return UserFollowersResponse{ResponseMeta: meta}, err
	}

	return UserFollowersResponse{
		Data:         parsed.JSON200.Data,
		ResponseMeta: meta,
	}, nil
}

// List retrieves all available seasons on the HackTheBox platform.
// This returns a comprehensive list of all seasons, including current and past seasons.
//
// Example:
//
//	seasons, err := client.Seasons.List(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, season := range seasons.Data {
//		fmt.Printf("Season: %s (ID: %d)\n", season.Name, season.Id)
//	}
func (s *Service) List(ctx context.Context) (ListResponse, error) {
	resp, err := s.base.Client.V4().GetSeasonList(s.base.Client.Limiter().Wrap(ctx))
	if err != nil {
		return ListResponse{ResponseMeta: common.ResponseMeta{}}, err
	}

	parsed, meta, err := common.Parse(resp, v4Client.ParseGetSeasonListResponse)
	if err != nil {
		return ListResponse{ResponseMeta: meta}, err
	}

	return ListResponse{
		Data:         parsed.JSON200.Data,
		ResponseMeta: meta,
	}, nil
}

// Machines retrieves all machines available in the current season.
// This returns information about machines that are part of the active season,
// including their difficulty, points, and availability status.
//
// Example:
//
//	machines, err := client.Seasons.Machines(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, machine := range machines.Data {
//		fmt.Printf("Machine: %s (Difficulty: %s)\n", machine.Name, machine.Difficulty)
//	}
func (s *Service) Machines(ctx context.Context) (MachinesResponse, error) {
	resp, err := s.base.Client.V4().GetSeasonMachines(s.base.Client.Limiter().Wrap(ctx))
	if err != nil {
		return MachinesResponse{ResponseMeta: common.ResponseMeta{}}, err
	}

	parsed, meta, err := common.Parse(resp, v4Client.ParseGetSeasonMachinesResponse)
	if err != nil {
		return MachinesResponse{ResponseMeta: meta}, err
	}

	return MachinesResponse{
		Data:         parsed.JSON200.Data,
		ResponseMeta: meta,
	}, nil
}

// ActiveMachine retrieves information about the currently active machine in the season.
// This returns details about the machine that is currently available for solving
// in the active season.
//
// Example:
//
//	activeMachine, err := client.Seasons.ActiveMachine(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Active machine: %s (ID: %d)\n", activeMachine.Data.Name, activeMachine.Data.Id)
func (s *Service) ActiveMachine(ctx context.Context) (ActiveMachineResponse, error) {
	resp, err := s.base.Client.V4().GetSeasonMachineActive(s.base.Client.Limiter().Wrap(ctx))
	if err != nil {
		return ActiveMachineResponse{ResponseMeta: common.ResponseMeta{}}, err
	}

	parsed, meta, err := common.Parse(resp, v4Client.ParseGetSeasonMachineActiveResponse)
	if err != nil {
		return ActiveMachineResponse{ResponseMeta: meta}, err
	}

	return ActiveMachineResponse{
		Data:         parsed.JSON200.Data,
		ResponseMeta: meta,
	}, nil
}
