/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
import React from 'react'
import { Providers, ProviderLabels } from '@/data/Providers'
import { ReactComponent as NullProviderIcon } from '@/images/integrations/null.svg'

const NullProvider = {
  id: Providers.NULL, // Unique ID, for a Provider (alphanumeric, lowercase)
  enabled: false, // Enabled Flag
  multiConnection: false, // If Provider is Multi-connection
  name: ProviderLabels.NULL, // Display Name of Data Provider
  // eslint-disable-next-line max-len
  icon: <NullProviderIcon className='providerIconSvg' width='30' height='30' style={{ float: 'left', marginTop: '5px' }} />, // Provider Icon
  iconDashboard: <NullProviderIcon className='providerIconSvg' width='48' height='48' />, // Provider Icon on INTEGRATIONS Dashboard
  settings: ({ activeProvider, activeConnection, isSaving, setSettings }) => (<></>) // REACT Settings Component for Render
}

export {
  NullProvider
}
