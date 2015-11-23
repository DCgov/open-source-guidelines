---
layout: page
title:  "Working in the Open"
---

Working in the open means providing public access to your source code and draft material throughout development.  Working in the open leads to a better end-product by encouraging collaboration within and outside of DC government and ensuring public feedback can be incorporated from day one.

> #### A word on security
A common concern about open source and working in the open is security.  As DC government moves toward a "default open" policy for its data, code, and content its stance is that openness fundamentally improves security by [making bugs shallow](https://en.wikipedia.org/wiki/Linus%27s_Law) and encouraging strong coding practices.  Given that a few key conditions are met, agencies should consider open sourcing their code and working in the open.  One of the main goals of the current pilot is to create a process that maximizes security and minimizes cost for open source efforts.

## Open Source Checklists 
Below are common scenarios teams face when publishing their work and working in the open.  Use the relevant checklist to evaluate your project. If you and your manager agree that the answer to each question is "Yep!," it's appropriate to immediately publish your work and continue development in the open.    

If you have doubts regarding any checklist item it's still very likely you can proceed.  Reach out to [opensource@dc.gov](mailto:opensource@dc.gov) for guidance. 

---

### Checklist: I'm using GitHub to work on a new project in the open
Your team is kicking off a new project and wants to develop in the open.  Make sure that:

- The repository's contents will be publicly releasable under an open license.
- The project does not have a special categorical security sensivity such as including classified information.
- The project or program manager agrees that the above conditions are met.

---

### Checklist: I'm going to publish an existing DC Government project to GitHub
Your team wants to make an existing codebase public for the first time.  Make sure that:

- All personally identifiable information (PII), including private keys, tokens, passwords, IP addresses, and other authentication information has been removed from the codebase.
- The licenses and contracts governing the project's contents permit their public release.
- I have included a README file with a clear project description. If applicable these include installation instructions.
- I have included a [LICENSE file]({{ site.baseurl }}/licensing) that describes how the project is licensed and notes the presence of third-party code with different license requirements.
- The project or program manager has ensured that the above steps are followed.

---

### Checklist: I'm forking a third-party GitHub repository for use in a DC Government project
Your team wants to reuse non-DC government open source code.  Make sure that:

- I understand the terms under which the source repository is licensed.  
- I have determined that the source repository's license permits modification and re-use.

---

### Checklist: I'm pushing or pulling code to an existing DC Government GitHub repository
As a routine part of working in the open your team members will make incremental changes to your public open source code. Team members should make sure that:

- All personally identifiable information (PII), including private keys, tokens, passwords, IP addresses, and other authentication information has been removed.
- All of the commits' contents are allowed to be released under the repository's existing license.

---

### Checklist: I'm using GitHub to work on a new project in a private repository

Creating private repositories within DCgov is discouraged. If you believe you have a valid reason to do so, please send an email explaining the repository's proposed contents and your requirement to [opensource@dc.gov](mailto:opensource@dc.gov).

---