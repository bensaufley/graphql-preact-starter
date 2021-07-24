import { FunctionComponent, h, JSX } from 'preact';
import { useCallback, useState } from 'preact/hooks';

import {
  HomeSubscription,
  useAddTodoMutation,
  useDeleteTodoMutation,
  useHomeSubscription,
} from '~components/Home/home.generated';

const ListItem = ({
  id,
  contents,
  status,
}: HomeSubscription['watchTodos'] extends Array<infer A> ? A : never) => {
  const [_, deleteMutation] = useDeleteTodoMutation();
  const deleteTodo: JSX.MouseEventHandler<HTMLButtonElement> = useCallback(
    (e) => {
      e.preventDefault();
      deleteMutation({ id: id.toString() });
    },
    [deleteMutation, id],
  );

  return (
    <li>
      {id}. {contents} ({status}){' '}
      <button type="button" onClick={deleteTodo}>
        &times;
      </button>
    </li>
  );
};

const Home: FunctionComponent = () => {
  const [{ data, error }] = useHomeSubscription();
  const [_, addTodo] = useAddTodoMutation();

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
      {error && <p>Error: {error.message}</p>}
      {!error && (
        <ul>
          {data?.watchTodos.map(({ id, contents, status }) => (
            <ListItem id={id} contents={contents} status={status} />
          ))}
        </ul>
      )}
    </div>
  );
};

export default Home;
