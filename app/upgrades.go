package app

import (
	"fmt"

	sdkmath "cosmossdk.io/math"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v5/modules/apps/29-fee/types"
)

func (app *App) RegisterUpgradeHandlers(experimental bool) {
	upgradeHandlerV1 := func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		m, err := app.mm.RunMigrations(ctx, app.configurator, fromVM)
		if err != nil {
			return m, err
		}
		// clear extra_eips from evm parameters
		// Ref: https://github.com/crypto-org-chain/cronos/issues/755
		params := app.EvmKeeper.GetParams(ctx)
		params.ExtraEIPs = []int64{}

		// fix the incorrect value on testnet parameters
		zero := sdkmath.ZeroInt()
		params.ChainConfig.LondonBlock = &zero

		app.EvmKeeper.SetParams(ctx, params)
		return m, nil
	}
	// `v1.0.0` upgrade plan will clear the `extra_eips` parameters, and upgrade ibc-go to v5.2.0.
	planName := "v1.0.0"
	app.UpgradeKeeper.SetUpgradeHandler(planName, upgradeHandlerV1)
	// "v1.0.0-testnet3-2" is another coordinated upgrade on testnet3 to upgrade ibc-go to "v5.2.0".
	planNameTestnet3 := "v1.0.0-testnet3-2"
	app.UpgradeKeeper.SetUpgradeHandler(planNameTestnet3, func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		return app.mm.RunMigrations(ctx, app.configurator, fromVM)
	})

	gravityPlanName := "v0.8.0-gravity-alpha3"
	if experimental {
		app.UpgradeKeeper.SetUpgradeHandler(gravityPlanName, func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
			updatedVM, err := app.mm.RunMigrations(ctx, app.configurator, fromVM)
			if err != nil {
				return nil, err
			}
			// set new gravity id
			gravParams := app.GravityKeeper.GetParams(ctx)
			gravParams.GravityId = "cronos_gravity_pioneer_v4"
			app.GravityKeeper.SetParams(ctx, gravParams)

			// Estimate time upgrade take place
			// 100% is not necessary here because it will be tuned by relayer later on
			// it is set to georli height at Mon Oct 31 2022 03:38:08 GMT+0900
			app.GravityKeeper.MigrateGravityContract(
				ctx, "0x0000000000000000000000000000000000000000", 7863000)

			return updatedVM, nil
		})
	}

	upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		if upgradeInfo.Name == planName {
			storeUpgrades := storetypes.StoreUpgrades{
				Added: []string{ibcfeetypes.StoreKey},
			}

			// configure store loader that checks if version == upgradeHeight and applies store upgrades
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
		}
	}
}
