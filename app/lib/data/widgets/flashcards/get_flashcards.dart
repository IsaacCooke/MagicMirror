import 'package:flutter/cupertino.dart';

import 'package:graphql_flutter/graphql_flutter.dart';

import 'package:app/data/models/constants.dart';
import 'package:app/data/models/flashcard.dart';

class GetFlashcards extends StatelessWidget {
  GetFlashcards({Key? key}) : super(key: key);

  static const String query = """
    query {
      getAllFlashcards {
        ID
        Term
        Definition
      }
    }
  """;

  final Constants constants = Constants();

  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      navigationBar: const CupertinoNavigationBar(
        middle: Text("Flashcards"),
      ),
      child: GraphQLProvider(
        client: constants.client,
        child: Query(
          options: QueryOptions(
            document: gql(query),
          ),
          builder: (result, {fetchMore, refetch}){
            if (result.isLoading){
              return const Center(
                child: CupertinoActivityIndicator(),
              );
            }
            if (result.data == null){
              return const Center(
                child: Text("No Data"),
              );
            }
            final flashcards = result.data!["getAllFlashcards"] as List<dynamic>;
            return ListView.builder(
              itemCount: flashcards.length,
              itemBuilder: (context, index){
                final flashcard = flashcards[index];
                final int id = flashcard["ID"];
                final String term = flashcard["Term"];
                final String definition = flashcard["Definition"];
                return FlashcardWidget(
                    id: id,
                    term: term,
                    definition: definition
                );
              },
            );
          },
        ),
      ),
    );
  }
}