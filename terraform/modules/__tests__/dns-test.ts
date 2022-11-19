// Copyright (c) HashiCorp, Inc
// SPDX-License-Identifier: MPL-2.0
import "cdktf/lib/testing/adapters/jest"; // Load types for expect matchers
import { TerraformStack, Testing } from "cdktf";
import { HappyDNS } from "../dns/main"
import { Route53Record } from "@cdktf/provider-aws/lib/route53-record";
import { AwsProvider } from "@cdktf/provider-aws/lib/provider";
describe("happy DNS", () => {
  it("creates a AWS Route53 record", () => {
    const construct = Testing.synthScope((scope) => {
      new HappyDNS(scope, "myDNS", "albName", "albZoneId", "zone.czi.technology", "zoneId");
    })
    expect(construct).toHaveResourceWithProperties(Route53Record, {
      type: "A",
      name: "stackName-appName.zone.czi.technology",
      zone_id: "zoneId", // weird, it changes from camelCase to snake_case otherwise the test will fail
      alias: [{
        name: "albName",
        zone_id: "albZoneId",
        evaluate_target_health: false,
      }]
    });
  });

  describe("can synthesize so", () => {
    const app = Testing.app();
    const stack = new TerraformStack(app, "test-happy-dns");
    new HappyDNS(stack, "my-dns", "albName", "albZoneId", "zone.czi.technology", "zoneId");
    new AwsProvider(stack, "AWS", {
      region: "us-west-2",
    });
    const synth = Testing.synth(stack)
    const fullSynth = Testing.fullSynth(stack)

    it("it looks like it did before", () => {
      expect(synth).toMatchSnapshot();
    });

    it("is valid terraform", () => {
      expect(fullSynth).toBeValidTerraform();
    });

    // TODO: need a way to configure the provider or something, not sure how to get this to pass
    it("plans without error", () => {
      expect(fullSynth).toPlanSuccessfully()
    });
  });
})
