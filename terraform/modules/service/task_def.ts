import { CloudwatchLogGroup } from "@cdktf/provider-aws/lib/cloudwatch-log-group";
import { EcsTaskDefinition } from "@cdktf/provider-aws/lib/ecs-task-definition";
import { IamRole } from "@cdktf/provider-aws/lib/iam-role"
import { DataAwsIamPolicyDocument } from "@cdktf/provider-aws/lib/data-aws-iam-policy-document";
import { Pet } from "@cdktf/provider-random/lib/pet";

import { Construct } from "constructs";
import { makeName } from "./main";
import { ContainerDefinition, HappyServiceMeta } from "./types";

export interface HappyECSTaskDefinitionProps {
    meta: HappyServiceMeta,
    executionRoleArn: string,
    tags?: { [key: string]: string };
}

export class HappyECSTaskDefinition extends Construct {
    readonly taskDef: EcsTaskDefinition
    readonly containerDefs: ContainerDefinition[]

    constructor(
        scope: Construct,
        id: string,
        config: HappyECSTaskDefinitionProps,
    ) {
        super(scope, id)

        const logGroup = new CloudwatchLogGroup(scope, `cwg`, {
            tags: config.tags,
            retentionInDays: 365,
            name: `/happy/${config.meta.env}/${config.meta.stackName}/${config.meta.serviceDef.name}`
        })

        this.containerDefs = [{
            name: config.meta.serviceDef.name,
            memory: config.meta.serviceDef.computeLimits.mem,
            image: config.meta.serviceDef.image,
            essential: true,
            environment: config.meta.serviceDef.environment,
            portMappings: [{ containerPort: config.meta.serviceDef.port }],
            logConfiguration: {
                logDriver: 'awslogs',
                options: {
                    "awslogs-group": logGroup.id,
                    "awslogs-region": config.meta.region,
                    "awslogs-stream-prefix": config.meta.env,
                }
            },
        }]

        const taskRolePolicyDoc = new DataAwsIamPolicyDocument(scope, `taskrolepolicy`, {
            statement: [{
                principals: [{
                    type: "Service",
                    identifiers: ["ecs-tasks.amazonaws.com"],
                }],
                actions: ["sts:AssumeRole"],
            }],
        })

        const pet = new Pet(scope, "pettask", { prefix: config.meta.stackName })
        const taskRole = new IamRole(scope, `iamrole`, {
            tags: config.tags,
            name: pet.id,
            assumeRolePolicy: taskRolePolicyDoc.json,
        })

        this.taskDef = new EcsTaskDefinition(scope, `ecstd`, {
            family: makeName(config.meta),
            memory: config.meta.serviceDef.computeLimits.mem,
            cpu: config.meta.serviceDef.computeLimits.cpu,
            networkMode: "awsvpc",
            requiresCompatibilities: ["FARGATE"],
            taskRoleArn: taskRole.arn,
            executionRoleArn: config.executionRoleArn,
            containerDefinitions: JSON.stringify(this.containerDefs),
        })
    }
}
