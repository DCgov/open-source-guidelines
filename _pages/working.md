---
layout: page
title: "Working in the Open"
---

Working in the open means providing public access to your source code throughout its development. This leads to a better end-product by encouraging collaboration within and outside of DC government and allowing public feedback from day one.

> #### A Word on Security
Security is a commonly voiced concern about open development. As DC government moves toward an "open by default" policy for its data, code, and content, it maintains that openness fundamentally improves security by [making bugs shallow](https://en.wikipedia.org/wiki/Linus%27s_Law) and encouraging strong coding practices. So long as a few key conditions are met, agencies should consider open sourcing their code and working in the open.

## Open Source Checklists

Below are common scenarios teams may face as part of the Open Source Pilot. Use the relevant checklist to evaluate your situation. If you and your manager can say "Yep!" to each item, you should go ahead with your plans.

Even if you have doubts regarding a checklist item, it is still very likely you can proceed. Reach out to [opensource@dc.gov](mailto:opensource@dc.gov) for guidance.

---

### Checklist: I am using GitHub to work on a new project in the open

My team is kicking off a new project and wants to develop in the open. I have ensured that:

- The repository's contents will be publicly releasable under an [open license]({{ site.baseurl }}/licensing/).
- The project does not have a special categorical security sensivity, such as including classified information.
- The project or program manager agrees that the above conditions have been met.

---

### Checklist: I am going to publish an existing DC Government project to GitHub

My team wants to make an existing codebase public for the first time. I have ensured that:

- All personally identifiable information (PII), including private keys, tokens, passwords, IP addresses, and other authentication information has been removed from the codebase.
- The licenses and contracts governing the project's contents permit their public release.
- The repository includes a [README file]({{ site.baseurl }}/documenting) with a clear project description. If applicable, these include installation instructions.
- The repository includes a [LICENSE file]({{ site.baseurl }}/licensing), which also notes the presence of any third-party code with different license requirements.
- The project or program manager agrees that the above steps have been followed.

---

### Checklist: I am forking a third-party GitHub repository for use in a DC Government project

My team wants to reuse non-DC Government open source code. I have ensured that:

- I understand the terms under which the source repository is licensed.
- The source repository's license permits modification and re-use.

---

### Checklist: I am pushing or pulling code to an existing DC Government GitHub repository

I want to make incremental changes to my team's public repository—a routine part of working in the open. I have ensured that:

- All personally identifiable information (PII), including private keys, tokens, passwords, IP addresses, and other authentication information has been removed.
- All of the commits' contents are allowed to be released under the repository's existing license.

---

### Checklist: I am using GitHub to work on a new project in a private repository

Creating private repositories within DCgov is discouraged. If you believe you have a valid reason to do so, please send an email explaining the repository's proposed contents and your requirement to [opensource@dc.gov](mailto:opensource@dc.gov).
