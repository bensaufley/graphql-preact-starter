import '~components/App/index.css';

import { Client, createClient, Provider as UrqlProvider } from '@urql/preact';
import { FunctionComponent, h } from 'preact';

import styles from '~components/App/styles.modules.css';
import Home from '~components/Home';

const App: FunctionComponent<{ client?: Client }> = ({
  client = createClient({ url: '/graphql' }),
}) => (
  <UrqlProvider value={client}>
    <header class={styles.siteHeader}>
      <h1>GraphQL App</h1>
    </header>
    <nav />
    <main>
      <Home />
    </main>
  </UrqlProvider>
);

export default App;
