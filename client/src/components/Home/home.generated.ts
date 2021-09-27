// Generated file. Do not edit!

import * as Urql from '@urql/preact';
import gql from 'graphql-tag';

import type * as SchemaTypes from '~graphql/schema.generated';

export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
export type GetTodosQueryVariables = SchemaTypes.Exact<{ [key: string]: never }>;

export type GetTodosQuery = {
  __typename?: 'Query';
  todos: Array<{
    __typename?: 'Todo';
    id: string;
    contents: string;
    status: SchemaTypes.TodoStatus;
  }>;
};

export type TodoAddedSubscriptionVariables = SchemaTypes.Exact<{ [key: string]: never }>;

export type TodoAddedSubscription = {
  __typename?: 'Subscription';
  todo: { __typename?: 'Todo'; id: string; contents: string; status: SchemaTypes.TodoStatus };
};

export type TodoUpdatedSubscriptionVariables = SchemaTypes.Exact<{ [key: string]: never }>;

export type TodoUpdatedSubscription = {
  __typename?: 'Subscription';
  todo: { __typename?: 'Todo'; id: string; contents: string; status: SchemaTypes.TodoStatus };
};

export type TodoDeletedSubscriptionVariables = SchemaTypes.Exact<{ [key: string]: never }>;

export type TodoDeletedSubscription = { __typename?: 'Subscription'; todoID: string };

export type AddTodoMutationVariables = SchemaTypes.Exact<{
  contents: SchemaTypes.Scalars['String'];
}>;

export type AddTodoMutation = {
  __typename?: 'Mutation';
  addTodo: SchemaTypes.Maybe<{ __typename?: 'Todo'; id: string }>;
};

export type AdvanceTodoMutationVariables = SchemaTypes.Exact<{
  id: SchemaTypes.Scalars['ID'];
}>;

export type AdvanceTodoMutation = {
  __typename?: 'Mutation';
  advanceTodo: SchemaTypes.Maybe<{ __typename?: 'Todo'; id: string }>;
};

export type TransitionTodoMutationVariables = SchemaTypes.Exact<{
  id: SchemaTypes.Scalars['ID'];
  status: SchemaTypes.TodoStatus;
}>;

export type TransitionTodoMutation = {
  __typename?: 'Mutation';
  transitionTodo: SchemaTypes.Maybe<{
    __typename?: 'Todo';
    id: string;
    status: SchemaTypes.TodoStatus;
  }>;
};

export type DeleteTodoMutationVariables = SchemaTypes.Exact<{
  id: SchemaTypes.Scalars['ID'];
}>;

export type DeleteTodoMutation = { __typename?: 'Mutation'; deleteTodo: boolean };

export const GetTodosDocument = gql`
  query GetTodos {
    todos: getTodos {
      id
      contents
      status
    }
  }
`;

export function useGetTodosQuery(
  options: Omit<Urql.UseQueryArgs<GetTodosQueryVariables>, 'query'> = {},
) {
  return Urql.useQuery<GetTodosQuery>({ query: GetTodosDocument, ...options });
}
export const TodoAddedDocument = gql`
  subscription TodoAdded {
    todo: todoAdded {
      id
      contents
      status
    }
  }
`;

export function useTodoAddedSubscription<TData = TodoAddedSubscription>(
  options: Omit<Urql.UseSubscriptionArgs<TodoAddedSubscriptionVariables>, 'query'> = {},
  handler?: Urql.SubscriptionHandler<TodoAddedSubscription, TData>,
) {
  return Urql.useSubscription<TodoAddedSubscription, TData, TodoAddedSubscriptionVariables>(
    { query: TodoAddedDocument, ...options },
    handler,
  );
}
export const TodoUpdatedDocument = gql`
  subscription TodoUpdated {
    todo: todoUpdated {
      id
      contents
      status
    }
  }
`;

export function useTodoUpdatedSubscription<TData = TodoUpdatedSubscription>(
  options: Omit<Urql.UseSubscriptionArgs<TodoUpdatedSubscriptionVariables>, 'query'> = {},
  handler?: Urql.SubscriptionHandler<TodoUpdatedSubscription, TData>,
) {
  return Urql.useSubscription<TodoUpdatedSubscription, TData, TodoUpdatedSubscriptionVariables>(
    { query: TodoUpdatedDocument, ...options },
    handler,
  );
}
export const TodoDeletedDocument = gql`
  subscription TodoDeleted {
    todoID: todoDeleted
  }
`;

export function useTodoDeletedSubscription<TData = TodoDeletedSubscription>(
  options: Omit<Urql.UseSubscriptionArgs<TodoDeletedSubscriptionVariables>, 'query'> = {},
  handler?: Urql.SubscriptionHandler<TodoDeletedSubscription, TData>,
) {
  return Urql.useSubscription<TodoDeletedSubscription, TData, TodoDeletedSubscriptionVariables>(
    { query: TodoDeletedDocument, ...options },
    handler,
  );
}
export const AddTodoDocument = gql`
  mutation AddTodo($contents: String!) {
    addTodo(input: { contents: $contents }) {
      id
    }
  }
`;

export function useAddTodoMutation() {
  return Urql.useMutation<AddTodoMutation, AddTodoMutationVariables>(AddTodoDocument);
}
export const AdvanceTodoDocument = gql`
  mutation AdvanceTodo($id: ID!) {
    advanceTodo(id: $id) {
      id
    }
  }
`;

export function useAdvanceTodoMutation() {
  return Urql.useMutation<AdvanceTodoMutation, AdvanceTodoMutationVariables>(AdvanceTodoDocument);
}
export const TransitionTodoDocument = gql`
  mutation TransitionTodo($id: ID!, $status: TodoStatus!) {
    transitionTodo(id: $id, status: $status) {
      id
      status
    }
  }
`;

export function useTransitionTodoMutation() {
  return Urql.useMutation<TransitionTodoMutation, TransitionTodoMutationVariables>(
    TransitionTodoDocument,
  );
}
export const DeleteTodoDocument = gql`
  mutation DeleteTodo($id: ID!) {
    deleteTodo(id: $id)
  }
`;

export function useDeleteTodoMutation() {
  return Urql.useMutation<DeleteTodoMutation, DeleteTodoMutationVariables>(DeleteTodoDocument);
}
