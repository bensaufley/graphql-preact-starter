// Generated file. Do not edit!

import * as Urql from '@urql/preact';
import gql from 'graphql-tag';

import type * as SchemaTypes from '~graphql/schema.generated';

export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
export type HomeSubscriptionVariables = SchemaTypes.Exact<{ [key: string]: never }>;

export type HomeSubscription = {
  __typename?: 'Subscription';
  watchTodos: Array<{
    __typename?: 'Todo';
    id: string;
    contents: string;
    status: SchemaTypes.TodoStatus;
  }>;
};

export type AddTodoMutationVariables = SchemaTypes.Exact<{
  contents: SchemaTypes.Scalars['String'];
}>;

export type AddTodoMutation = {
  __typename?: 'Mutation';
  addTodo: SchemaTypes.Maybe<{ __typename?: 'Todo'; id: string }>;
};

export type DeleteTodoMutationVariables = SchemaTypes.Exact<{
  id: SchemaTypes.Scalars['ID'];
}>;

export type DeleteTodoMutation = { __typename?: 'Mutation'; deleteTodo: boolean };

export const HomeDocument = gql`
  subscription Home {
    watchTodos {
      id
      contents
      status
    }
  }
`;

export function useHomeSubscription<TData = HomeSubscription>(
  options: Omit<Urql.UseSubscriptionArgs<HomeSubscriptionVariables>, 'query'> = {},
  handler?: Urql.SubscriptionHandler<HomeSubscription, TData>,
) {
  return Urql.useSubscription<HomeSubscription, TData, HomeSubscriptionVariables>(
    { query: HomeDocument, ...options },
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
export const DeleteTodoDocument = gql`
  mutation DeleteTodo($id: ID!) {
    deleteTodo(id: $id)
  }
`;

export function useDeleteTodoMutation() {
  return Urql.useMutation<DeleteTodoMutation, DeleteTodoMutationVariables>(DeleteTodoDocument);
}
