/*
 * SPDX-License-Identifier: Apache-2.0
 * Copyright 2018-2020 The Feast Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package feast.ingestion.transform.fn;

import feast.proto.types.FeatureRowProto.FeatureRow;
import org.apache.beam.sdk.transforms.DoFn;

public class ProcessFeatureRowDoFn extends DoFn<FeatureRow, FeatureRow> {

  private String defaultProject;

  public ProcessFeatureRowDoFn(String defaultProject) {
    this.defaultProject = defaultProject;
  }

  @ProcessElement
  public void processElement(ProcessContext context) {
    FeatureRow featureRow = context.element();
    String featureSetId = stripVersion(featureRow.getFeatureSet());
    featureSetId = applyDefaultProject(featureSetId);
    featureRow = featureRow.toBuilder().setFeatureSet(featureSetId).build();
    context.output(featureRow);
  }

  // For backward compatibility. Will be deprecated eventually.
  private String stripVersion(String featureSetId) {
    String[] split = featureSetId.split(":");
    return split[0];
  }

  private String applyDefaultProject(String featureSetId) {
    String[] split = featureSetId.split("/");
    if (split.length == 1) {
      return defaultProject + "/" + featureSetId;
    }
    return featureSetId;
  }
}
