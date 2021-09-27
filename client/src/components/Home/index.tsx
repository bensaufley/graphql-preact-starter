import { FunctionComponent, h, JSX } from 'preact';
import { useCallback, useMemo, useState } from 'preact/hooks';

import {
  useAddTodoMutation,
  useAdvanceTodoMutation,
  useDeleteTodoMutation,
  useGetTodosQuery,
  useTodoAddedSubscription,
  useTodoDeletedSubscription,
  useTodoUpdatedSubscription,
  useTransitionTodoMutation,
} from '~components/Home/home.generated';
import { Todo, TodoStatus } from '~graphql/schema.generated';

const advanceableStatuses: TodoStatus[] = [TodoStatus.Unstarted, TodoStatus.InProgress];

const ListItem = ({ id, contents, status }: Todo) => {
  const [, advanceMutation] = useAdvanceTodoMutation();
  const [, transitionMutation] = useTransitionTodoMutation();
  const [, deleteMutation] = useDeleteTodoMutation();
  const isAdvanceable = useMemo(() => advanceableStatuses.includes(status), [status]);
  const advanceTodo: JSX.MouseEventHandler<HTMLButtonElement> = useCallback(
    (e) => {
      e.preventDefault();
      if (!isAdvanceable) return;

      advanceMutation({ id });
    },
    [advanceMutation, id, isAdvanceable],
  );
  const transitionTodo: JSX.GenericEventHandler<HTMLSelectElement> = useCallback(
    ({ currentTarget }) => {
      transitionMutation({ id, status: currentTarget!.value as TodoStatus });
    },
    [transitionMutation, id],
  );
  const deleteTodo: JSX.MouseEventHandler<HTMLButtonElement> = useCallback(
    (e) => {
      e.preventDefault();
      deleteMutation({ id });
    },
    [deleteMutation, id],
  );

  return (
    <li id={id} title={id}>
      {contents} (
      <select onChange={transitionTodo} value={status}>
        {Object.values(TodoStatus).map((st) => (
          <option value={st}>{st}</option>
        ))}
      </select>
      ){' '}
      {isAdvanceable && (
        <button type="button" onClick={advanceTodo}>
          &rarr;
        </button>
      )}{' '}
      <button type="button" onClick={deleteTodo}>
        &times;
      </button>
    </li>
  );
};

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
