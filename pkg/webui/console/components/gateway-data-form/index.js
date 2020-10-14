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

import React from 'react'
import bind from 'autobind-decorator'

import delay from '@console/constants/delays'

import Form from '@ttn-lw/components/form'
import Input from '@ttn-lw/components/input'
import Checkbox from '@ttn-lw/components/checkbox'
import SubmitBar from '@ttn-lw/components/submit-bar'
import UnitInput from '@ttn-lw/components/unit-input'
import KeyValueMap from '@ttn-lw/components/key-value-map'
import Collapse from '@ttn-lw/components/collapse'

import Message from '@ttn-lw/lib/components/message'

import { GsFrequencyPlansSelect } from '@console/containers/freq-plans-select'
import OwnersSelect from '@console/containers/owners-select'

import PropTypes from '@ttn-lw/lib/prop-types'
import sharedMessages from '@ttn-lw/lib/shared-messages'

import { unit as unitRegexp, emptyDuration as emptyDurationRegexp } from '@console/lib/regexp'

import m from './messages'
import validationSchema from './validation-schema'

class GatewayDataForm extends React.Component {
  static propTypes = {
    /** The SubmitBar content. */
    children: PropTypes.node.isRequired,
    collapse: PropTypes.bool,
    error: PropTypes.error,
    /** React reference to be passed to the form. */
    formRef: PropTypes.shape({}),
    initialValues: PropTypes.gateway,
    onSubmit: PropTypes.func.isRequired,
    update: PropTypes.bool,
  }

  static defaultProps = {
    collapse: false,
    formRef: undefined,
    update: false,
    error: '',
    initialValues: validationSchema.cast({}),
  }

  constructor(props) {
    super(props)

    this.state = {
      shouldDisplayWarning: this.isNotValidDuration(props.initialValues.schedule_anytime_delay),
    }
  }

  @bind
  onScheduleAnytimeDelayChange(value) {
    this.setState({ shouldDisplayWarning: this.isNotValidDuration(value) })
  }

  @bind
  onSubmit(values, helpers) {
    const { onSubmit } = this.props
    const castedValues = validationSchema.cast(values)

    onSubmit(castedValues, helpers)
  }

  decodeDelayValue(value) {
    if (emptyDurationRegexp.test(value)) {
      return {
        duration: undefined,
        unit: value,
      }
    }
    const duration = value.split(unitRegexp)[0]
    const unit = value.split(duration)[1]
    return {
      duration: duration ? Number(duration) : undefined,
      unit,
    }
  }

  isNotValidDuration(value) {
    const { duration, unit } = this.decodeDelayValue(value)
    switch (unit) {
      case 'ms':
        return duration < delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY
      case 's':
        return duration < delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY / 1000
      case 'm':
        return duration < delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY / 60000
      case 'h':
        return duration < delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY / 3600000
    }
  }

  render() {
    const { update, error, initialValues, formRef, children, collapse } = this.props
    const { shouldDisplayWarning } = this.state

    return (
      <Form
        error={error}
        onSubmit={this.onSubmit}
        initialValues={initialValues}
        validationSchema={validationSchema}
        formikRef={formRef}
      >
        <Collapse
          title={m.basicTitle}
          description={m.basicDescription}
          disabled={false}
          initialCollapsed={false}
        >
          <Message component="h4" content={sharedMessages.generalSettings} />
          {!update && <OwnersSelect name="owner_id" required autoFocus />}
          <Form.Field
            title={sharedMessages.gatewayID}
            name="ids.gateway_id"
            placeholder={m.gatewayIdPlaceholder}
            required
            disabled={update}
            component={Input}
          />
          <Form.Field
            title={sharedMessages.gatewayEUI}
            name="ids.eui"
            type="byte"
            min={8}
            max={8}
            placeholder={sharedMessages.gatewayEUI}
            component={Input}
          />
          <Form.Field
            title={sharedMessages.gatewayName}
            placeholder={m.gatewayNamePlaceholder}
            name="name"
            component={Input}
          />
          <Form.Field
            title={sharedMessages.gatewayDescription}
            description={m.gatewayDescDescription}
            placeholder={m.gatewayDescPlaceholder}
            name="description"
            type="textarea"
            component={Input}
          />
          <Form.Field
            title={sharedMessages.gatewayServerAddress}
            description={m.gsServerAddressDescription}
            placeholder={sharedMessages.addressPlaceholder}
            name="gateway_server_address"
            component={Input}
          />
          <Form.Field
            title={sharedMessages.gatewayStatus}
            name="status_public"
            component={Checkbox}
            label={sharedMessages.public}
            description={m.statusDescription}
          />
          <Form.Field
            name="attributes"
            title={sharedMessages.attributes}
            keyPlaceholder={sharedMessages.key}
            valuePlaceholder={sharedMessages.value}
            addMessage={sharedMessages.addAttributes}
            component={KeyValueMap}
            description={sharedMessages.attributeDescription}
          />
          <Message component="h4" content={sharedMessages.gatewayUpdateOptions} />
          <Form.Field
            title={sharedMessages.automaticUpdates}
            name="auto_update"
            component={Checkbox}
            description={m.autoUpdateDescription}
          />
          <Form.Field
            title={sharedMessages.channel}
            description={m.updateChannelDescription}
            placeholder={sharedMessages.stable}
            name="update_channel"
            component={Input}
            autoComplete="on"
          />
        </Collapse>
        <Collapse
          title={m.lorawanTitle}
          description={m.lorawanDescription}
          disabled={false}
          initialCollapsed={collapse}
        >
          <GsFrequencyPlansSelect name="frequency_plan_id" menuPlacement="top" />
          <Form.Field
            title={sharedMessages.gatewayScheduleDownlinkLate}
            name="schedule_downlink_late"
            component={Checkbox}
            description={m.scheduleDownlinkLateDescription}
          />
          <Form.Field
            title={m.dutyCycle}
            name="enforce_duty_cycle"
            component={Checkbox}
            label={m.enforced}
            description={m.enforceDutyCycleDescription}
          />
          <Form.Field
            title={m.scheduleAnyTimeDelay}
            name="schedule_anytime_delay"
            component={UnitInput}
            description={{
              ...m.scheduleAnyTimeDescription,
              values: {
                minimumValue: delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY,
                defaultValue: delay.DEFAULT_GATEWAY_SCHEDULE_ANYTIME_DELAY,
              },
            }}
            units={[
              { label: sharedMessages.milliseconds, value: 'ms' },
              { label: sharedMessages.seconds, value: 's' },
              { label: sharedMessages.minutes, value: 'm' },
              { label: sharedMessages.hours, value: 'h' },
            ]}
            onChange={this.onScheduleAnytimeDelayChange}
            warning={
              shouldDisplayWarning
                ? {
                    ...m.delayWarning,
                    values: { minimumValue: delay.MINIMUM_GATEWAY_SCHEDULE_ANYTIME_DELAY },
                  }
                : undefined
            }
            required
          />
        </Collapse>
        <SubmitBar>{children}</SubmitBar>
      </Form>
    )
  }
}

export default GatewayDataForm
