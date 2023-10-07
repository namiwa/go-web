---
path: getting-started-with-miniconda
date: 2023-06-30T00:00:00.856Z
title: Getting Started with Python & Miniconda
category: coding
---

### Python

[Python](https://www.python.org/about/) has been around for over 30 years, and remains one of the most popular
languages loved by businesses, researchers and hobbyist alike. It has a small,
simple syntax which easy to pick up by beginners, and extendible enough to create
tools tackling different domains, from battle-tested production websites to its
solid ecosystem of libraries and frameworks in machine learning and data science.

Python is a great language to begin with for people who are starting out in their
programming journey. Personally, while the native Python downloads for Windows, MacOS
and Linux are sufficient for most applications, I think that using a dedicated Python
package manager will make it much easier to install specific binary tools and
packages which would otherwise require involved methods (like building such packages
from scratch) which would hamper one's speed in learning or using the language.

### R

[R](https://www.r-project.org) was built with data science in mind. It function
is to make data analytics easier, and let users visualize their data transformations
easier and be able to run their analysis in real-time in the command line.
While my experience had only been limited to tidyverse and its related libraries,
it has a concise syntax where chaining operations on data is done in a much neater
way than in Python.

### Motivation

Given my exposure to Python is more towards development work rather then data science
related and I had only briefly fiddled around with R, I do feel like both of these
languages shine when used in the right development environment. Which is why, if
possible, I would use a package manager to get started immediately with both languages.

### Introducing Miniconda

[Anaconda](http://anaconda.com) simplifies Python package management, where it not only
manages different virtual environments for Python, it also manages all the necessary
binaries that each package need as best as possible. What's more, this tool is
platform agnostic, so as long at the operating system is supported, it should work
and run the same across every platform!

From my experience, using the full version of Anaconda along with its Graphical User
Interface (GUI) felt quite bloated, despite being able to meet all of my development needs.
Fortunately, they also release and maintain a stripped down Command Line Interface (CLI)
version where it only contains the core package manager with just the bare essential packages,
called [Miniconda](https://docs.conda.io/en/latest/miniconda.html).

### Installing Miniconda

First, download the latest shell installer for Miniconda [here.](https://docs.conda.io/en/latest/miniconda.html#latest-miniconda-installer-links)
Follow the given instructions to verify that the script downloaded matches the hash published.

After running the relevant `conda init` command for the terminal shell of choice
of your platform, you should be greeted with the following prompt in your terminal:

![conda base environment example](../images/getting-started-with-miniconda/miniconda-base.png)

The `(base)` next to the shell prompt indicates that miniconda is operating in its
default environment and that the initalization had been successful. Another way to
check would be to run `conda --version` which would give the following output:

![conda version](../images/getting-started-with-miniconda/miniconda-version.png)

#### Creating Python environment

To get started with setting up a clean Python conda environment, run the following
commands in the terminal:

- `conda create -n python-3-10 python=3.10`
  - Note that you are able to use any name, `python-3-10` is chosen for this example
- Enter yes prompt

![conda create python env](../images/getting-started-with-miniconda/miniconda-create-python-env.png)

- `conda activate python-3-10`
  ![conda activate python env](../images/getting-started-with-miniconda/miniconda-activate-python-env.png)

Running `python` or `which python` helpes to determine if the miniconda managed
python had been installed correctly:

![conda check python env](../images/getting-started-with-miniconda/miniconda-check-python-env.png)

#### Creating R environment

To use R from miniconda, first exit out of any conda managed python environment
by running `conda deactivate`

- Create empty environment by running `conda create --name r_env`:
  ![conda create r environment](../images/getting-started-with-miniconda/miniconda-create-r.png)

- Activate said environment (`conda activate r_env`) and install relevant R dependancies by running `conda install r r-essentials --channel conda-forge` and accepting downloads
  ![conda after install r environment](../images/getting-started-with-miniconda/miniconda-after-install-r.png)

- The install from the previous command will take a couple minutes, so hang tight!

- To test if R install is correct, run the following `R --version`: ![conda r version](../images/getting-started-with-miniconda/miniconda-r-version.png)


### Conclusion

With that, it should be possible to get started on any development work related to either
Python or R from miniconda! And the best part, this should work across major operating systems,
Windows, MacOS & Linux :D
