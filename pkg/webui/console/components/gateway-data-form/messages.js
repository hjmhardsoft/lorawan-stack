// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import { defineMessages } from 'react-intl'

const messages = defineMessages({
  enforced: 'Enforced',
  delayWarning:
    'Delay too short. The lower bound ({minimumValue}ms) will be used by the Gateway Server.',
  dutyCycle: 'Duty cycle',
  gatewayIdPlaceholder: 'my-new-gateway',
  gatewayNamePlaceholder: 'My new gateway',
  gsServerAddressDescription: 'The address of the Gateway Server to connect to',
  gatewayDescPlaceholder: 'Description for my new gateway',
  gatewayDescDescription:
    'Optional gateway description; can also be used to save notes about the gateway',
  statusDescription: 'The status of this gateway may be publicly displayed',
  scheduleDownlinkLateDescription: 'Enable server-side buffer of downlink messages',
  autoUpdateDescription: 'Gateway can be updated automatically',
  updateChannelDescription: 'Channel for gateway automatic updates',
  enforceDutyCycleDescription:
    'Recommended for all gateways in order to respect spectrum regulations',
  scheduleAnyTimeDelay: 'Schedule any time delay',
  scheduleAnyTimeDescription:
    'Configure gateway delay (minimum: {minimumValue}ms, default: {defaultValue}ms)',
})

export default messages
