import 'package:flutter/cupertino.dart';
import 'package:graphql_flutter/graphql_flutter.dart';

final HttpLink httpLink = HttpLink('https://8c26-80-195-185-6.ngrok-free.app/graphql');

class Constants {
  final ValueNotifier<GraphQLClient> client = ValueNotifier(
    GraphQLClient(
      cache: GraphQLCache(),
      link: httpLink,
    ),
  );
}