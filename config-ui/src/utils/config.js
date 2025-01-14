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

export const DEVLAKE_ENDPOINT = '/api'
export const GRAFANA_BASE_URL = '/grafana'
export const LOCAL_BASE_URL = 'http://localhost:3002'
export const GRAFANA_ENDPOINT = '/d/0Rjxknc7z/demo-homepage?orgId=1'
export const GRAFANA_URL = process.env.LOCAL ? LOCAL_BASE_URL + GRAFANA_ENDPOINT : GRAFANA_BASE_URL + GRAFANA_ENDPOINT
