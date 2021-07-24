import '~components/App/index.css';

import {
  Client,
  createClient,
  defaultExchanges,
  Provider as UrqlProvider,
  subscriptionExchange,
} from '@urql/preact';
import { FunctionComponent, h } from 'preact';
import { SubscriptionClient } from 'subscriptions-transport-ws';

import styles from '~components/App/styles.modules.css';
import Home from '~components/Home';

const App: FunctionComponent<{ client?: Client; subscriptionClient?: SubscriptionClient }> = ({
  subscriptionClient = new SubscriptionClient(
    `ws://${window.location.hostname}${
      window.location.port ? `:${window.location.port}` : ''
    }/graphql`,
    { reconnect: true, timeout: 10_000 },
  ),
  client = createClient({
    url: '/graphql',
    exchanges: [
      ...defaultExchanges,
      subscriptionExchange({
        forwardSubscription(operation) {
          return subscriptionClient.request(operation);
        },
      }),
    ],
  }),
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
