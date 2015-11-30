# DC Open Source Guidelines

These are draft guidelines for using and working on open source code in DC Government. They are being developed as part of a pilot program launched by the Office of Technology and Innovation in the [Office of the Chief Technology Officer](http://octo.dc.gov).

This is a living document, which we are using in our own work. We welcome collaboration, participation, and feedback from teams across DC government as we scale the program. 

## To run the site locally

Fork and clone the repo:

```
git clone git@github.com:DCgov/open-source-guidelines.git
cd open-source-guidelines
```

Ensure that you have Ruby (> version 2.0) installed. You can check your version by entering `ruby -v` at the command prompt.

Next, install the site's dependencies with Bundler:

```
gem install bundler
bundle install
```

And run the site locally:

```
./go serve
```

Open it up in your browser: <http://localhost:4000/>


## License

The project is in the public domain, and all contributions to it will be released as such. By submitting a pull request, you are agreeing to waive all rights to your contribution under the terms of the [CC0 Public Domain Dedication](http://creativecommons.org/publicdomain/zero/1.0/).

If you contribute the open source work of others, please mark it clearly in your pull request.
