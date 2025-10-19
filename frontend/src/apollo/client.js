import { ApolloClient, InMemoryCache, createHttpLink } from '@apollo/client';

// Configure HTTP Link para conectar com o BFF GraphQL
const httpLink = createHttpLink({
  uri: import.meta.env.PROD 
    ? 'http://localhost:8080/graphql' // Em produção, conecta ao BFF na rede do Docker
    : 'http://localhost:8080/graphql', // Em desenvolvimento, conecta localmente
  credentials: 'same-origin', // Include cookies if needed
});

// Configure Apollo Client com cache e link
const client = new ApolloClient({
  link: httpLink,
  cache: new InMemoryCache({
    typePolicies: {
      Query: {
        fields: {
          orders: {
            // Cache policy para orders - sempre buscar dados frescos
            fetchPolicy: 'cache-and-network',
          },
          payments: {
            // Cache policy para payments
            fetchPolicy: 'cache-and-network',
          },
          notifications: {
            // Cache policy para notifications
            fetchPolicy: 'cache-and-network',
          },
        },
      },
    },
  }),
  defaultOptions: {
    watchQuery: {
      // Poll a cada 5 segundos por padrão para dados em tempo real
      pollInterval: 5000,
      errorPolicy: 'all',
    },
    query: {
      errorPolicy: 'all',
    },
  },
});

export default client;