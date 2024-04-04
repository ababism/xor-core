FROM python:3.12-slim-bullseye as prod


RUN pip install poetry==1.4.2

# Configuring poetry
RUN poetry config virtualenvs.create false

# Copying requirements of a project

WORKDIR /app/src

COPY --chown=application:application ../pyproject.toml ./
COPY --chown=application:application ../poetry.lock ./

RUN apt-get update && \
    apt-get --yes --no-install-recommends install make="4.3-4.1" && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Installing requirements
RUN poetry install

# Copying actuall application
COPY . /app/src/
RUN poetry install

CMD ["/usr/local/bin/python", "-m", "payments"]

FROM prod as dev

RUN poetry install
