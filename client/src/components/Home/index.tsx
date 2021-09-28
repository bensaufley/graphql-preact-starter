import { FunctionComponent, h, JSX } from 'preact';
import { useCallback, useState } from 'preact/hooks';

import {
  useAddTodoMutation,
  useGetTodosQuery,
  useTodoAddedSubscription,
  useTodoDeletedSubscription,
  useTodoUpdatedSubscription,
} from '~components/Home/home.generated';
import ListItem from '~components/ListItem';

const Home: FunctionComponent = () => {
  const [{ data, error, fetching }] = useGetTodosQuery();
  const [_, addTodo] = useAddTodoMutation();
  useTodoAddedSubscription();
  useTodoUpdatedSubscription();
  useTodoDeletedSubscription();

  const [newTodoContent, setNewTodoContent] = useState('');
  const handleChange: JSX.GenericEventHandler<HTMLInputElement> = useCallback(
    ({ currentTarget }) => {
      setNewTodoContent(currentTarget.value);
    },
    [setNewTodoContent],
  );

  const createTodo = useCallback(
    (e: Event) => {
      e.preventDefault();
      addTodo({ contents: newTodoContent });
      setNewTodoContent('');
    },
    [addTodo, newTodoContent],
  );

  return (
    <div>
      <h2>Home</h2>
      <form onSubmit={createTodo}>
        <input type="text" onChange={handleChange} name="contents" value={newTodoContent} />
        <button type="submit">Create</button>
      </form>
      {fetching && <p>Loading&hellip;</p>}
      {!fetching && error && <p>Error: {error.message}</p>}
      {!fetching && !error && (
        <ul>
          {data?.todos?.map(({ id, contents, status }) => (
            <ListItem id={id} contents={contents} status={status} />
          ))}
        </ul>
      )}
    </div>
  );
};

export default Home;
