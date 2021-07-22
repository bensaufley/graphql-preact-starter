// Generated file. Do not edit!

import * as Urql from '@urql/preact';
import gql from 'graphql-tag';

import type * as SchemaTypes from '~graphql/schema.generated';

export type Omit<T, K extends keyof T> = Pick<T, Exclude<keyof T, K>>;
export type HomeQueryVariables = SchemaTypes.Exact<{ [key: string]: never }>;

export type HomeQuery = {
  __typename?: 'Query';
  foo: SchemaTypes.Maybe<{ __typename?: 'Bar'; name: string }>;
};

export const HomeDocument = gql`
  query Home {
    foo {
      name
    }
  }
`;

export function useHomeQuery(options: Omit<Urql.UseQueryArgs<HomeQueryVariables>, 'query'> = {}) {
  return Urql.useQuery<HomeQuery>({ query: HomeDocument, ...options });
}
