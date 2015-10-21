---
layout: page
title:  "Working in the Open"
---

The District is moving toward a policy of "open by default." This means that unless there is a compelling reason to do otherwise, we intend to:

* Use Free and Open Source Software (FOSS), which is software that does not charge users a purchase or licensing fee for modifying or redistributing the source code, in our projects and contribute back to the open source community.
* Develop our work in the open.
* Publish publicly all source code created or modified by DC Government agencies, whether developed in-house by government staff or through contracts negotiated by those agencies.

Currently, abiding by the second two points requires going through a rigorous security and privacy review process.

This page lays out several scenarios and provides a 'fast track' checklist to help your team accelerate the review process in cases where the security risk is low. If you and/or your supervisor have doubts regarding any checklist item, it is likely a signal that a more extensive review may be warranted.

On a final note, this document assumes basic fluency in key Git concepts such as fork, and repository. If you or your team need orientation on Git, Github, or Open Source generally, a great overview can be found [here](https://18f.gsa.gov/2015/03/03/how-to-use-github-and-the-terminal-a-guide/). You may also arrange a training session by sending an email to opensource@dc.gov.

### Scenario 1: I want to fork a third-party GitHub repository for use in a DC Government project

- I understand the terms under which the source repository is licensed.
- I have determined that the source repository's license permits modification and re-use.

### Scenario 2: I want to publish an existing DC Government project to GitHub

- I have removed all private keys, tokens, passwords, and other authentication information.
- I have removed all personally identifiable information (PII).
- I have determined that the license(s) and contract(s) governing the project's contents permit their public release.
- I have sent an email to opensource@dc.gov, laying out the repository's proposed content and received confirmation.

### Scenario 3: I want to use GitHub to work on a new project in the open

- I have confirmed that the repository's contents will be publicly releasable under an open license.
- I have determined that the project poses no undue risk to DC Government staff or infrastructure, or to District residents.
- My supervisor has reviewed my analysis and agreed in writing (email) that I can move forward.
- I have sent an email to opensource@dc.gov laying out the repository's proposed content and received confirmation.

### Scenario 4: I want to push or pull commits to an existing DCgov GitHub repository

- My commits do not contain any private keys, tokens, passwords, or other authentication information.
- My commits do not contain any personally identifiable information.
- My commits' contents are in line with the repository's existing license.
