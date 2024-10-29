// Copyright 2024 aivruu
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the
// Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package main

// AttributesContainer This struct is used as container for all additional repository's details.
type AttributesContainer struct {
	forked   bool    // Whether repository is forked.
	parent   *string // Repository's original owner's name (if repo is forked).
	canFork  bool    // Whether this repository can be forked.
	stars    int32
	forks    int32
	private  bool
	archived bool
	disabled bool
	language string // Repository's main language.
	topics   *[]string
}

func NewAttributesContainer(forked bool, parent *string, canFork bool, stars, forks int32, private bool, archived bool, disabled bool, language string, topics *[]string) AttributesContainer {
	return AttributesContainer{
		forked:   forked,
		parent:   parent,
		canFork:  canFork,
		stars:    stars,
		forks:    forks,
		private:  private,
		archived: archived,
		disabled: disabled,
		language: language,
		topics:   topics,
	}
}

// Forked This method returns whether the repository is a fork.
func (r *AttributesContainer) Forked() bool {
	return r.forked
}

// Parent This method returns the original owner's name of the repository (if the repository is a fork).
func (r *AttributesContainer) Parent() *string {
	return r.parent
}

// CanFork This method returns whether the repository can be forked.
func (r *AttributesContainer) CanFork() bool {
	return r.canFork
}

// Stars This method returns the number of stars the repository has.
func (r *AttributesContainer) Stars() int32 {
	return r.stars
}

// Forks This method returns the number of forks the repository has.
func (r *AttributesContainer) Forks() int32 {
	return r.forks
}

// Public This method returns whether the repository is public.
func (r *AttributesContainer) Public() bool {
	return !r.private
}

// Archived This method returns whether the repository is archived.
func (r *AttributesContainer) Archived() bool {
	return r.archived
}

// Disabled This method returns whether the repository is disabled.
func (r *AttributesContainer) Disabled() bool {
	return r.disabled
}

// Language This method returns the main language used for the repository.
func (r *AttributesContainer) Language() string {
	return r.language
}

// Topics This method returns the topics of the repository, or nil if isn't available.
func (r *AttributesContainer) Topics() *[]string {
	return r.topics
}
