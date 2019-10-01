// +build !ignore_autogenerated

/*
Copyright 2019 kubeflow.org.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	constants "github.com/kubeflow/kfserving/pkg/constants"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlibiExplainerSpec) DeepCopyInto(out *AlibiExplainerSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlibiExplainerSpec.
func (in *AlibiExplainerSpec) DeepCopy() *AlibiExplainerSpec {
	if in == nil {
		return nil
	}
	out := new(AlibiExplainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomSpec) DeepCopyInto(out *CustomSpec) {
	*out = *in
	in.Container.DeepCopyInto(&out.Container)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomSpec.
func (in *CustomSpec) DeepCopy() *CustomSpec {
	if in == nil {
		return nil
	}
	out := new(CustomSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSpec) DeepCopyInto(out *DeploymentSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSpec.
func (in *DeploymentSpec) DeepCopy() *DeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointSpec) DeepCopyInto(out *EndpointSpec) {
	*out = *in
	in.Predictor.DeepCopyInto(&out.Predictor)
	if in.Explainer != nil {
		in, out := &in.Explainer, &out.Explainer
		*out = new(ExplainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Transformer != nil {
		in, out := &in.Transformer, &out.Transformer
		*out = new(TransformerSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointSpec.
func (in *EndpointSpec) DeepCopy() *EndpointSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in EndpointStatusMap) DeepCopyInto(out *EndpointStatusMap) {
	{
		in := &in
		*out = make(EndpointStatusMap, len(*in))
		for key, val := range *in {
			var outVal *StatusConfigurationSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(StatusConfigurationSpec)
				**out = **in
			}
			(*out)[key] = outVal
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointStatusMap.
func (in EndpointStatusMap) DeepCopy() EndpointStatusMap {
	if in == nil {
		return nil
	}
	out := new(EndpointStatusMap)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExplainerConfig) DeepCopyInto(out *ExplainerConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExplainerConfig.
func (in *ExplainerConfig) DeepCopy() *ExplainerConfig {
	if in == nil {
		return nil
	}
	out := new(ExplainerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExplainerSpec) DeepCopyInto(out *ExplainerSpec) {
	*out = *in
	if in.Alibi != nil {
		in, out := &in.Alibi, &out.Alibi
		*out = new(AlibiExplainerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomSpec)
		(*in).DeepCopyInto(*out)
	}
	out.DeploymentSpec = in.DeploymentSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExplainerSpec.
func (in *ExplainerSpec) DeepCopy() *ExplainerSpec {
	if in == nil {
		return nil
	}
	out := new(ExplainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExplainersConfig) DeepCopyInto(out *ExplainersConfig) {
	*out = *in
	out.AlibiExplainer = in.AlibiExplainer
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExplainersConfig.
func (in *ExplainersConfig) DeepCopy() *ExplainersConfig {
	if in == nil {
		return nil
	}
	out := new(ExplainersConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeastTransformerSpec) DeepCopyInto(out *FeastTransformerSpec) {
	*out = *in
	if in.EntityIds != nil {
		in, out := &in.EntityIds, &out.EntityIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FeatureIds != nil {
		in, out := &in.FeatureIds, &out.FeatureIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeastTransformerSpec.
func (in *FeastTransformerSpec) DeepCopy() *FeastTransformerSpec {
	if in == nil {
		return nil
	}
	out := new(FeastTransformerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KFService) DeepCopyInto(out *KFService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KFService.
func (in *KFService) DeepCopy() *KFService {
	if in == nil {
		return nil
	}
	out := new(KFService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KFService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KFServiceList) DeepCopyInto(out *KFServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KFService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KFServiceList.
func (in *KFServiceList) DeepCopy() *KFServiceList {
	if in == nil {
		return nil
	}
	out := new(KFServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KFServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KFServiceSpec) DeepCopyInto(out *KFServiceSpec) {
	*out = *in
	in.Default.DeepCopyInto(&out.Default)
	if in.Canary != nil {
		in, out := &in.Canary, &out.Canary
		*out = new(EndpointSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KFServiceSpec.
func (in *KFServiceSpec) DeepCopy() *KFServiceSpec {
	if in == nil {
		return nil
	}
	out := new(KFServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KFServiceStatus) DeepCopyInto(out *KFServiceStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.Default != nil {
		in, out := &in.Default, &out.Default
		*out = new(EndpointStatusMap)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[constants.KFServiceEndpoint]*StatusConfigurationSpec, len(*in))
			for key, val := range *in {
				var outVal *StatusConfigurationSpec
				if val == nil {
					(*out)[key] = nil
				} else {
					in, out := &val, &outVal
					*out = new(StatusConfigurationSpec)
					**out = **in
				}
				(*out)[key] = outVal
			}
		}
	}
	if in.Canary != nil {
		in, out := &in.Canary, &out.Canary
		*out = new(EndpointStatusMap)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[constants.KFServiceEndpoint]*StatusConfigurationSpec, len(*in))
			for key, val := range *in {
				var outVal *StatusConfigurationSpec
				if val == nil {
					(*out)[key] = nil
				} else {
					in, out := &val, &outVal
					*out = new(StatusConfigurationSpec)
					**out = **in
				}
				(*out)[key] = outVal
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KFServiceStatus.
func (in *KFServiceStatus) DeepCopy() *KFServiceStatus {
	if in == nil {
		return nil
	}
	out := new(KFServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ONNXSpec) DeepCopyInto(out *ONNXSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ONNXSpec.
func (in *ONNXSpec) DeepCopy() *ONNXSpec {
	if in == nil {
		return nil
	}
	out := new(ONNXSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorConfig) DeepCopyInto(out *PredictorConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorConfig.
func (in *PredictorConfig) DeepCopy() *PredictorConfig {
	if in == nil {
		return nil
	}
	out := new(PredictorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorSpec) DeepCopyInto(out *PredictorSpec) {
	*out = *in
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Tensorflow != nil {
		in, out := &in.Tensorflow, &out.Tensorflow
		*out = new(TensorflowSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TensorRT != nil {
		in, out := &in.TensorRT, &out.TensorRT
		*out = new(TensorRTSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.XGBoost != nil {
		in, out := &in.XGBoost, &out.XGBoost
		*out = new(XGBoostSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SKLearn != nil {
		in, out := &in.SKLearn, &out.SKLearn
		*out = new(SKLearnSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ONNX != nil {
		in, out := &in.ONNX, &out.ONNX
		*out = new(ONNXSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PyTorch != nil {
		in, out := &in.PyTorch, &out.PyTorch
		*out = new(PyTorchSpec)
		(*in).DeepCopyInto(*out)
	}
	out.DeploymentSpec = in.DeploymentSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorSpec.
func (in *PredictorSpec) DeepCopy() *PredictorSpec {
	if in == nil {
		return nil
	}
	out := new(PredictorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PredictorsConfig) DeepCopyInto(out *PredictorsConfig) {
	*out = *in
	out.Tensorflow = in.Tensorflow
	out.TensorRT = in.TensorRT
	out.Xgboost = in.Xgboost
	out.SKlearn = in.SKlearn
	out.PyTorch = in.PyTorch
	out.ONNX = in.ONNX
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PredictorsConfig.
func (in *PredictorsConfig) DeepCopy() *PredictorsConfig {
	if in == nil {
		return nil
	}
	out := new(PredictorsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PyTorchSpec) DeepCopyInto(out *PyTorchSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PyTorchSpec.
func (in *PyTorchSpec) DeepCopy() *PyTorchSpec {
	if in == nil {
		return nil
	}
	out := new(PyTorchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SKLearnSpec) DeepCopyInto(out *SKLearnSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SKLearnSpec.
func (in *SKLearnSpec) DeepCopy() *SKLearnSpec {
	if in == nil {
		return nil
	}
	out := new(SKLearnSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusConfigurationSpec) DeepCopyInto(out *StatusConfigurationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusConfigurationSpec.
func (in *StatusConfigurationSpec) DeepCopy() *StatusConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(StatusConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TensorRTSpec) DeepCopyInto(out *TensorRTSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TensorRTSpec.
func (in *TensorRTSpec) DeepCopy() *TensorRTSpec {
	if in == nil {
		return nil
	}
	out := new(TensorRTSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TensorflowSpec) DeepCopyInto(out *TensorflowSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TensorflowSpec.
func (in *TensorflowSpec) DeepCopy() *TensorflowSpec {
	if in == nil {
		return nil
	}
	out := new(TensorflowSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TransformerSpec) DeepCopyInto(out *TransformerSpec) {
	*out = *in
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = new(CustomSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Feast != nil {
		in, out := &in.Feast, &out.Feast
		*out = new(FeastTransformerSpec)
		(*in).DeepCopyInto(*out)
	}
	out.DeploymentSpec = in.DeploymentSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TransformerSpec.
func (in *TransformerSpec) DeepCopy() *TransformerSpec {
	if in == nil {
		return nil
	}
	out := new(TransformerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *XGBoostSpec) DeepCopyInto(out *XGBoostSpec) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new XGBoostSpec.
func (in *XGBoostSpec) DeepCopy() *XGBoostSpec {
	if in == nil {
		return nil
	}
	out := new(XGBoostSpec)
	in.DeepCopyInto(out)
	return out
}
