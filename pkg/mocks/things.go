// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package mocks

import (
	"context"
	"strconv"
	"sync"

	"github.com/MainfluxLabs/mainflux"
	"github.com/MainfluxLabs/mainflux/pkg/errors"
	"github.com/MainfluxLabs/mainflux/things"
)

var _ things.Service = (*mainfluxThings)(nil)

type mainfluxThings struct {
	mu          sync.Mutex
	counter     uint64
	things      map[string]things.Thing
	channels    map[string]things.Channel
	auth        mainflux.AuthServiceClient
	connections map[string][]string
}

// NewThingsService returns Mainflux Things service mock.
// Only methods used by SDK are mocked.
func NewThingsService(things map[string]things.Thing, channels map[string]things.Channel, auth mainflux.AuthServiceClient) things.Service {
	return &mainfluxThings{
		things:      things,
		channels:    channels,
		auth:        auth,
		connections: make(map[string][]string),
	}
}

func (svc *mainfluxThings) CreateThings(_ context.Context, owner string, ths ...things.Thing) ([]things.Thing, error) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	userID, err := svc.auth.Identify(context.Background(), &mainflux.Token{Value: owner})
	if err != nil {
		return []things.Thing{}, errors.ErrAuthentication
	}
	for i := range ths {
		svc.counter++
		ths[i].Owner = userID.Email
		ths[i].ID = strconv.FormatUint(svc.counter, 10)
		ths[i].Key = ths[i].ID
		svc.things[ths[i].ID] = ths[i]
	}

	return ths, nil
}

func (svc *mainfluxThings) ViewThing(_ context.Context, owner, id string) (things.Thing, error) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	userID, err := svc.auth.Identify(context.Background(), &mainflux.Token{Value: owner})
	if err != nil {
		return things.Thing{}, errors.ErrAuthentication
	}

	if t, ok := svc.things[id]; ok && t.Owner == userID.Email {
		return t, nil

	}

	return things.Thing{}, errors.ErrNotFound
}

func (svc *mainfluxThings) Connect(_ context.Context, owner, chID string, thIDs []string) error {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	userID, err := svc.auth.Identify(context.Background(), &mainflux.Token{Value: owner})
	if err != nil {
		return errors.ErrAuthentication
	}

	if svc.channels[chID].Owner != userID.Email {
		return errors.ErrAuthentication
	}
	svc.connections[chID] = append(svc.connections[chID], thIDs...)

	return nil
}

func (svc *mainfluxThings) Disconnect(_ context.Context, owner, chID string, thIDs []string) error {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	userID, err := svc.auth.Identify(context.Background(), &mainflux.Token{Value: owner})
	if err != nil {
		return errors.ErrAuthentication
	}

	if svc.channels[chID].Owner != userID.Email {
		return errors.ErrAuthentication
	}

	ids := svc.connections[chID]
	var count int
	var newConns []string
	for _, thID := range thIDs {
		for _, id := range ids {
			if id == thID {
				count++
				continue
			}
			newConns = append(newConns, id)
		}

		if len(newConns)-len(ids) != count {
			return errors.ErrNotFound
		}
		svc.connections[chID] = newConns
	}

	return nil
}

func (svc *mainfluxThings) ListThingsByIDs(ctx context.Context, thingIDs []string) (things.Page, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) RemoveThings(_ context.Context, owner string, ids ...string) error {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	userID, err := svc.auth.Identify(context.Background(), &mainflux.Token{Value: owner})
	if err != nil {
		return errors.ErrAuthentication
	}

	for _, id := range ids {
		if t, ok := svc.things[id]; !ok || t.Owner != userID.Email {
			return errors.ErrNotFound
		}

		delete(svc.things, id)
	}

	return nil
}

func (svc *mainfluxThings) ViewChannel(_ context.Context, owner, id string) (things.Channel, error) {
	if c, ok := svc.channels[id]; ok {
		return c, nil
	}
	return things.Channel{}, errors.ErrNotFound
}

func (svc *mainfluxThings) UpdateThing(context.Context, string, things.Thing) error {
	panic("not implemented")
}

func (svc *mainfluxThings) UpdateKey(context.Context, string, string, string) error {
	panic("not implemented")
}

func (svc *mainfluxThings) ListThings(context.Context, string, bool, things.PageMetadata) (things.Page, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) ViewChannelByThing(context.Context, string, string) (things.Channel, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) ListThingsByChannel(context.Context, string, string, things.PageMetadata) (things.Page, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) Backup(context.Context, string) (things.Backup, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) Restore(context.Context, string, things.Backup) error {
	panic("not implemented")
}

func (svc *mainfluxThings) CreateChannels(_ context.Context, owner string, chs ...things.Channel) ([]things.Channel, error) {
	svc.mu.Lock()
	defer svc.mu.Unlock()

	userID, err := svc.auth.Identify(context.Background(), &mainflux.Token{Value: owner})
	if err != nil {
		return []things.Channel{}, errors.ErrAuthentication
	}
	for i := range chs {
		svc.counter++
		chs[i].Owner = userID.Email
		chs[i].ID = strconv.FormatUint(svc.counter, 10)
		svc.channels[chs[i].ID] = chs[i]
	}

	return chs, nil
}

func (svc *mainfluxThings) UpdateChannel(context.Context, string, things.Channel) error {
	panic("not implemented")
}

func (svc *mainfluxThings) ListChannels(context.Context, string, bool, things.PageMetadata) (things.ChannelsPage, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) RemoveChannels(context.Context, string, ...string) error {
	panic("not implemented")
}

func (svc *mainfluxThings) ViewChannelProfile(ctx context.Context, chID string) (things.Profile, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) GetConnByKey(context.Context, string) (things.Connection, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) IsChannelOwner(context.Context, string, string) error {
	panic("not implemented")
}

func (svc *mainfluxThings) Identify(context.Context, string) (string, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) ShareThing(ctx context.Context, token, thingID string, actions, userIDs []string) error {
	panic("not implemented")
}

func findIndex(list []string, val string) int {
	for i, v := range list {
		if v == val {
			return i
		}
	}

	return -1
}

func (svc *mainfluxThings) ListGroupThings(ctx context.Context, token, groupID string, pm things.PageMetadata) (things.GroupThingsPage, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) ListGroupThingsByChannel(ctx context.Context, token, grID, chID string, pm things.PageMetadata) (things.GroupThingsPage, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) CreateGroups(ctx context.Context, token string, groups ...things.Group) ([]things.Group, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) ListGroups(ctx context.Context, token string, admin bool, pm things.PageMetadata) (things.GroupPage, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) ListGroupsByIDs(ctx context.Context, groupIDs []string) ([]things.Group, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) RemoveGroups(ctx context.Context, token string, ids ...string) error {
	panic("not implemented")
}

func (svc *mainfluxThings) UpdateGroup(ctx context.Context, token string, group things.Group) (things.Group, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) ViewGroup(ctx context.Context, token, id string) (things.Group, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) AssignThing(ctx context.Context, token string, groupID string, thingIDs ...string) error {
	panic("not implemented")
}

func (svc *mainfluxThings) UnassignThing(ctx context.Context, token string, groupID string, thingIDs ...string) error {
	panic("not implemented")
}

func (svc *mainfluxThings) ViewThingMembership(ctx context.Context, token string, thingID string) (things.Group, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) AssignChannel(ctx context.Context, token string, groupID string, channelIDs ...string) error {
	panic("not implemented")
}

func (svc *mainfluxThings) UnassignChannel(ctx context.Context, token string, groupID string, channelIDs ...string) error {
	panic("not implemented")
}

func (svc *mainfluxThings) ViewChannelMembership(ctx context.Context, token string, channelID string) (things.Group, error) {
	panic("not implemented")
}

func (svc *mainfluxThings) ListGroupChannels(ctx context.Context, token, groupID string, pm things.PageMetadata) (things.GroupChannelsPage, error) {
	panic("not implemented")
}
