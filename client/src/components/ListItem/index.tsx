import { h, JSX } from 'preact';
import { useCallback, useMemo } from 'preact/hooks';

import {
  useAdvanceTodoMutation,
  useDeleteTodoMutation,
  useTransitionTodoMutation,
} from '~components/Home/home.generated';
import styles from '~components/ListItem/styles.modules.css';
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
      {contents}{' '}
      <select class={styles.select} onChange={transitionTodo} value={status}>
        {Object.keys(TodoStatus).map((st) => (
          <option value={(TodoStatus as any)[st as any]}>{st}</option>
        ))}
      </select>{' '}
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

export default ListItem;
