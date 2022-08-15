/*
 * Copyright 2022 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logger

// Logger logger root interface
type Logger interface {
	Debug(args ...any)
	Debugf(template string, args ...any)
	Info(args ...any)
	Infof(template string, args ...any)
	Infow(template string, args ...any)
	Warn(args ...any)
	Warnf(template string, args ...any)
	Warnw(template string, args ...any)
	Error(args ...any)
	Errorf(template string, args ...any)
	Errorw(template string, args ...any)
	Fatal(args ...any)
	Fatalf(template string, args ...any)
	Fatalw(template string, args ...any)
	Panic(args ...any)
	Panicf(template string, args ...any)
	Panicw(template string, args ...any)
	Sync()
}
