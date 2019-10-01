# Copyright 2019 kubeflow.org.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from setuptools import setup, find_packages

TESTS_REQUIRE = [
    'pytest',
    'pytest-tornasync',
    'mypy'
]

setup(
    name='feasttransformer',
    version='0.1.0',
    author_email='ellisbigelow@seldon.io',
    license='../../LICENSE.txt',
    url='https://github.com/kubeflow/kfserving/python/kfserving/feasttransformer',
    description='KFTransformer that reads from a Feast Feature Store to transform requests.',
    long_description=open('README.md').read(),
    python_requires='>=3.6',
    packages=find_packages("feasttransformer"),
    install_requires=[
        "kfserving>=0.1.0",
        "requests3>=0.0.0",
        "pandas>=0.24.2"
    ],
    tests_require=TESTS_REQUIRE,
    extras_require={'test': TESTS_REQUIRE}
)
