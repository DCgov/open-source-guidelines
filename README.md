# DC Open Source Guidelines

These are draft guidelines for using and working on open source code in DC Government. They are being developed as part of a pilot program launched by the Office of Technology and Innovation in the [Office of the Chief Technology Officer](http://octo.dc.gov).

This is a living document, which we are using in our own work. We welcome collaboration, participation, and feedback from teams across DC government as we scale the program. 

## To run the site locally

Fork and clone the repo:

```
$ git clone git@github.com:DCgov/open-source-guidelines.git
$ cd open-source-guidelines
```

Ensure that you have Ruby (> version 2.0) installed. You can check your version by entering `ruby -v` at the command prompt. You can install Ruby most quickly with Homebrew:

```
$ brew update
$ brew install ruby
```

Whether or not Ruby is already installed, we strongly recommend using a Ruby version manager such as [rbenv](https://github.com/sstephenson/rbenv) or [rvm](https://rvm.io/) to help ensure that Ruby version upgrades don't mean all your [gems](https://rubygems.org/) will need to be rebuilt.

Next, install the site's dependencies with Bundler:

```
$ gem install bundler
$ bundle install
```

And run the site locally:

```
$ ./go serve
```

Open it up in your browser: <http://localhost:4000/>


## Public Domain

This project is in the public domain within the United States, and copyright and related rights in the work worldwide are waived through the CC0 1.0 Universal public domain dedication. For more information, see [LICENSE.md](https://github.com/DCgov/license/blob/master/LICENSE.md).
