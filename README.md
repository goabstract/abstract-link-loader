# abstract-link-loader

## EXPERIMENTAL/WIP

This loader is not ready for production use and is still a WIP. PRs welcomed!

## Usage

Create a Abstract "link" file in your project:

```sh
$ create-react-app my-project
$ cd my-project
$ echo -n 'export default { url: "ABSTRACT_SHARE_URL_FOR_ICONS_FILE" }' > src/icons.link.js
```

Import styles from any Abstract layer:

```js
import { styles } from "./icons.link";
```

Start webpack with `ABSTRACT_TOKEN` defined in environment:

```sh
$ ABSTRACT_TOKEN=my-token npm start
```

Visit https://sdk.goabstract.com/docs/authentication/ to learn more about abstract tokens.
