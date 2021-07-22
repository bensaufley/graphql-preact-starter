import { FunctionComponent, h } from 'preact';

import { useHomeQuery } from '~components/Home/home.generated';

const Home: FunctionComponent = () => {
  const queryResponse = useHomeQuery();

  return (
    <div>
      <h2>Home</h2>
      <code>
        <pre>{queryResponse}</pre>
      </code>
    </div>
  );
};

export default Home;
