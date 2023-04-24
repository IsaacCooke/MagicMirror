import 'package:flutter/cupertino.dart';

import 'package:graphql_flutter/graphql_flutter.dart';

import 'package:app/data/models/constants.dart';
import 'package:app/data/models/flashcard.dart';

class DeleteFlashcards extends StatefulWidget {
  const DeleteFlashcards({Key? key}) : super(key: key);

  DeleteFlashcardsState createState() => DeleteFlashcardsState();
}

class DeleteFlashcardsState extends State<DeleteFlashcards> {

  static const String query = """
    query {
      getAllFlashcards {
        ID
        Term
        Definition
      }
    }
  """;

  void _deleteFlashcard(int id) async {
    final client = constants.client;

    const String mutation = """
    mutation deleteFlashcard(\$id: Int!) {
      deleteFlashcard(id: \$id) {
        ID
      }
    }
  """;

    final options = MutationOptions(
      document: gql(mutation),
      variables: {
        "id": id,
      },
    );

    final result = await client.value.mutate(options);

    if(result.hasException) {
      throw Exception(result.exception.toString());
    }
  }

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
                return Row(
                  mainAxisSize: MainAxisSize.min,
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Flexible(
                      child: FlashcardWidget(
                        id: id,
                        term: term,
                        definition: definition,
                      ),
                    ),
                    CupertinoButton(
                      child: const Icon(CupertinoIcons.delete),
                      onPressed: () {
                        _deleteFlashcard(id);
                        refetch!();
                      },
                    ),
                  ],
                );
              },
            );
          },
        ),
      ),
    );
  }
}