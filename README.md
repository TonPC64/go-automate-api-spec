# Demo for automate api spec by git action

> docs: https://tonpc64.github.io/go-automate-api-spec

![Go_ API Spec Automation](https://github.com/user-attachments/assets/795c3937-47c7-4e09-b1ff-4c0d852c50c8)

[Slide](https://docs.google.com/presentation/d/e/2PACX-1vRBu7Bg4NUbecP5hbtyzLNzQi2JA4YiE6RzcJczFL3jcqDjRbQfFqQugAhwKpLQ7Rj5cgcko4YWOtAC/pub?start=false&loop=false&delayms=3000)

## Setup and Usage

Follow these steps to set up and generate the API spec:

```bash
# Install pre-commit and swag
make prepare

# Install project dependencies
make install

# Generate the Swagger file
make generate.api-spec

# View the result of the API spec
make dev.document
```
