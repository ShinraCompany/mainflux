// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

//go:build !test

package api

import (
	"context"
	"time"

	"github.com/MainfluxLabs/mainflux/things"
	"github.com/go-kit/kit/metrics"
)

var _ things.Service = (*metricsMiddleware)(nil)

type metricsMiddleware struct {
	counter metrics.Counter
	latency metrics.Histogram
	svc     things.Service
}

// MetricsMiddleware instruments core service by tracking request count and latency.
func MetricsMiddleware(svc things.Service, counter metrics.Counter, latency metrics.Histogram) things.Service {
	return &metricsMiddleware{
		counter: counter,
		latency: latency,
		svc:     svc,
	}
}

func (ms *metricsMiddleware) CreateThings(ctx context.Context, token string, ths ...things.Thing) (saved []things.Thing, err error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "create_things").Add(1)
		ms.latency.With("method", "create_things").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.CreateThings(ctx, token, ths...)
}

func (ms *metricsMiddleware) UpdateThing(ctx context.Context, token string, thing things.Thing) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "update_thing").Add(1)
		ms.latency.With("method", "update_thing").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.UpdateThing(ctx, token, thing)
}

func (ms *metricsMiddleware) UpdateKey(ctx context.Context, token, id, key string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "update_key").Add(1)
		ms.latency.With("method", "update_key").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.UpdateKey(ctx, token, id, key)
}

func (ms *metricsMiddleware) ViewThing(ctx context.Context, token, id string) (things.Thing, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "view_thing").Add(1)
		ms.latency.With("method", "view_thing").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ViewThing(ctx, token, id)
}

func (ms *metricsMiddleware) ListThings(ctx context.Context, token string, admin bool, pm things.PageMetadata) (things.Page, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "list_things").Add(1)
		ms.latency.With("method", "list_things").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ListThings(ctx, token, admin, pm)
}

func (ms *metricsMiddleware) ListThingsByIDs(ctx context.Context, ids []string) (things.Page, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "list_things_by_ids").Add(1)
		ms.latency.With("method", "list_things_by_ids").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ListThingsByIDs(ctx, ids)
}

func (ms *metricsMiddleware) ListThingsByChannel(ctx context.Context, token, chID string, pm things.PageMetadata) (things.Page, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "list_things_by_channel").Add(1)
		ms.latency.With("method", "list_things_by_channel").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ListThingsByChannel(ctx, token, chID, pm)
}

func (ms *metricsMiddleware) RemoveThings(ctx context.Context, token string, id ...string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "remove_things").Add(1)
		ms.latency.With("method", "remove_things").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.RemoveThings(ctx, token, id...)
}

func (ms *metricsMiddleware) CreateChannels(ctx context.Context, token string, channels ...things.Channel) (saved []things.Channel, err error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "create_channels").Add(1)
		ms.latency.With("method", "create_channels").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.CreateChannels(ctx, token, channels...)
}

func (ms *metricsMiddleware) UpdateChannel(ctx context.Context, token string, channel things.Channel) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "update_channel").Add(1)
		ms.latency.With("method", "update_channel").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.UpdateChannel(ctx, token, channel)
}

func (ms *metricsMiddleware) ViewChannel(ctx context.Context, token, id string) (things.Channel, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "view_channel").Add(1)
		ms.latency.With("method", "view_channel").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ViewChannel(ctx, token, id)
}

func (ms *metricsMiddleware) ListChannels(ctx context.Context, token string, admin bool, pm things.PageMetadata) (things.ChannelsPage, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "list_channels").Add(1)
		ms.latency.With("method", "list_channels").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ListChannels(ctx, token, admin, pm)
}

func (ms *metricsMiddleware) ViewChannelByThing(ctx context.Context, token, thID string) (things.Channel, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "view_channel_by_thing").Add(1)
		ms.latency.With("method", "view_channel_by_thing").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ViewChannelByThing(ctx, token, thID)
}

func (ms *metricsMiddleware) RemoveChannels(ctx context.Context, token string, ids ...string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "remove_channels").Add(1)
		ms.latency.With("method", "remove_channels").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.RemoveChannels(ctx, token, ids...)
}

func (ms *metricsMiddleware) ViewChannelProfile(ctx context.Context, chID string) (things.Profile, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "view_channel_profile").Add(1)
		ms.latency.With("method", "view_channel_profile").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ViewChannelProfile(ctx, chID)
}

func (ms *metricsMiddleware) Connect(ctx context.Context, token, chID string, thIDs []string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "connect").Add(1)
		ms.latency.With("method", "connect").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.Connect(ctx, token, chID, thIDs)
}

func (ms *metricsMiddleware) Disconnect(ctx context.Context, token, chID string, thIDs []string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "disconnect").Add(1)
		ms.latency.With("method", "disconnect").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.Disconnect(ctx, token, chID, thIDs)
}

func (ms *metricsMiddleware) GetConnByKey(ctx context.Context, key string) (things.Connection, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "get_conn_by_key").Add(1)
		ms.latency.With("method", "get_conn_by_key").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.GetConnByKey(ctx, key)
}

func (ms *metricsMiddleware) IsChannelOwner(ctx context.Context, owner, chanID string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "is_channel_owner").Add(1)
		ms.latency.With("method", "is_channel_owner").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.IsChannelOwner(ctx, owner, chanID)
}

func (ms *metricsMiddleware) Identify(ctx context.Context, key string) (string, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "identify").Add(1)
		ms.latency.With("method", "identify").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.Identify(ctx, key)
}

func (ms *metricsMiddleware) Backup(ctx context.Context, token string) (bk things.Backup, err error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "backup").Add(1)
		ms.latency.With("method", "backup").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.Backup(ctx, token)
}

func (ms *metricsMiddleware) Restore(ctx context.Context, token string, backup things.Backup) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "restore").Add(1)
		ms.latency.With("method", "restore").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.Restore(ctx, token, backup)
}

func (ms *metricsMiddleware) CreateGroups(ctx context.Context, token string, grs ...things.Group) ([]things.Group, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "create_groups").Add(1)
		ms.latency.With("method", "create_groups").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.CreateGroups(ctx, token, grs...)
}

func (ms *metricsMiddleware) UpdateGroup(ctx context.Context, token string, g things.Group) (things.Group, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "update_group").Add(1)
		ms.latency.With("method", "update_group").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.UpdateGroup(ctx, token, g)
}

func (ms *metricsMiddleware) ViewGroup(ctx context.Context, token, id string) (things.Group, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "view_group").Add(1)
		ms.latency.With("method", "view_group").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ViewGroup(ctx, token, id)
}

func (ms *metricsMiddleware) ListGroups(ctx context.Context, token string, admin bool, pm things.PageMetadata) (things.GroupPage, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "list_group").Add(1)
		ms.latency.With("method", "list_group").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ListGroups(ctx, token, admin, pm)
}

func (ms *metricsMiddleware) ListGroupsByIDs(ctx context.Context, groupIDs []string) ([]things.Group, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "list_groups_by_ids").Add(1)
		ms.latency.With("method", "list_groups_by_ids").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ListGroupsByIDs(ctx, groupIDs)
}

func (ms *metricsMiddleware) ListGroupThings(ctx context.Context, token, groupID string, pm things.PageMetadata) (tp things.GroupThingsPage, err error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "list_group_things").Add(1)
		ms.latency.With("method", "list_group_things").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ListGroupThings(ctx, token, groupID, pm)
}

func (ms *metricsMiddleware) ListGroupThingsByChannel(ctx context.Context, token, grID, chID string, pm things.PageMetadata) (things.GroupThingsPage, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "list_group_things_by_channel").Add(1)
		ms.latency.With("method", "list_group_things_by_channel").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ListGroupThingsByChannel(ctx, token, grID, chID, pm)
}

func (ms *metricsMiddleware) ViewThingMembership(ctx context.Context, token, thingID string) (things.Group, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "view_thing_membership").Add(1)
		ms.latency.With("method", "view_thing_membership").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ViewThingMembership(ctx, token, thingID)
}

func (ms *metricsMiddleware) RemoveGroups(ctx context.Context, token string, ids ...string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "remove_groups").Add(1)
		ms.latency.With("method", "remove_groups").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.RemoveGroups(ctx, token, ids...)
}

func (ms *metricsMiddleware) AssignThing(ctx context.Context, token, groupID string, thingIDs ...string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "assign_thing").Add(1)
		ms.latency.With("method", "assign_thing").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.AssignThing(ctx, token, groupID, thingIDs...)
}

func (ms *metricsMiddleware) UnassignThing(ctx context.Context, token, groupID string, thingIDs ...string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "unassign_thing").Add(1)
		ms.latency.With("method", "unassign_thing").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.UnassignThing(ctx, token, groupID, thingIDs...)
}

func (ms *metricsMiddleware) AssignChannel(ctx context.Context, token, groupID string, channelIDs ...string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "assign_channel").Add(1)
		ms.latency.With("method", "assign_channel").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.AssignChannel(ctx, token, groupID, channelIDs...)
}

func (ms *metricsMiddleware) UnassignChannel(ctx context.Context, token, groupID string, channelIDs ...string) error {
	defer func(begin time.Time) {
		ms.counter.With("method", "unassign_channel").Add(1)
		ms.latency.With("method", "unassign_channel").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.UnassignChannel(ctx, token, groupID, channelIDs...)
}

func (ms *metricsMiddleware) ListGroupChannels(ctx context.Context, token, groupID string, pm things.PageMetadata) (things.GroupChannelsPage, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "list_group_channels").Add(1)
		ms.latency.With("method", "list_group_channels").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ListGroupChannels(ctx, token, groupID, pm)
}

func (ms *metricsMiddleware) ViewChannelMembership(ctx context.Context, token, channelID string) (things.Group, error) {
	defer func(begin time.Time) {
		ms.counter.With("method", "view_channel_membership").Add(1)
		ms.latency.With("method", "view_channel_membership").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return ms.svc.ViewChannelMembership(ctx, token, channelID)
}
