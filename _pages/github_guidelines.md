---
layout: page
title:  "Using Open Source"
---

Working in the open means providing public access to your source code and draft material - ideally from day one.  Working in the open leads to a better end-product by encouraging collaboration within and outside of DC government and ensuring public feedback can be incorporated from day one.

> #### A word on security
One of the most commonly voiced concerns for government about open source and working in the open is security.  As DC government moves toward a "default open" policy for its data, code, and content its stance is that openness fundamentally improves security by [making bugs shallow](https://en.wikipedia.org/wiki/Linus%27s_Law) and encouraging strong coding practices.  Given that a few key conditions are met, agencies should consider open sourcing their code and working in the open.  The current pilot is design, in part, to test light weight, repeatable that maximize security while minimizing cost and latency. 


## Open Source Scenarios 
Below are common scenarios teams face when considering publishing their work or working in the open. Based on your scenario, use the accompanying checklist to evaluate your project. If you and your supervisor agree that the answer to each question is "Yep!," it is appropriate to immediately publish your work and continue development in the open.    

If you have doubts regarding any checklist item it is still very likely you can proceed.  Reach out to [opensource@dc.gov](mailto:opensource@dc.gov) for guidance. 

### Scenario 1: I want to use GitHub to work on a new project in the open
In this scenario a new project is kicking off and the team wishes to develop in the open.  Verify the following:

- The repository's contents will be publicly releasable under an open license.
- The project does not have a special categorical security sensivity such as including classified information.
- The project or program manager agrees that the above conditions are met.


### Scenario 2: I want to publish an existing DC Government project to GitHub
In this scenario a team or agency wishes to publish an existing codebase for the first time.  Verify the following:

- All personally identifiable information (PII), including private keys, tokens, passwords, IP addresses, and other authentication information has been removed from the codebase.
- The licenses and contracts governing the project's contents permit their public release.
- I have included a README file with a clear project description. If applicable these include installation instructions.
- I have included a [LICENSE file]({{ site.baseurl }}/licensing) that describes how the project is licensed and notes the presence of third-party code with different license requirements.
- The project or program manager has ensured that the above steps are followed.


### Scenario 3: I want to fork a third-party GitHub repository for use in a DC Government project
In this scenario a DC government team wishs to make use of an existing open source project maintained by a third party.  Verify the following :

- I understand the terms under which the source repository is licensed.  
- I have determined that the source repository's license permits modification and re-use.


### Scenario 4: I want to use GitHub to work on a new project in a private repository

Creating private repositories within DCgov is discouraged. If you believe you have a valid reason to do so, please send an email explaining the repository's proposed contents and your requirement to [opensource@dc.gov](mailto:opensource@dc.gov).


### Scenario 5: I want to push or pull commits to an existing DCgov GitHub repository
In this scenario project members want to make incremental changes to an existing DC government open source project. This scenario is very common and constitutes much of the day to day work on an open source project.  Developers are expected to do this regularly as a routine part of their job, taking into account the following:

-  All personally identifiable information (PII), including private keys, tokens, passwords, IP addresses, and other authentication information has been removed.
- All of the commits' contents are allowed to be released under the repository's existing license.
