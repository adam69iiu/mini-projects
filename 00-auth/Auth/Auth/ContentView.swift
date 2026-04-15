//
//  ContentView.swift
//  Auth
//
//  Created by adam mohammed on 4/15/26.
//

import SwiftUI

struct ContentView: View {
    @State  var health: String = ""
    @State var errorMessage : String = ""
    var body: some View {
        VStack {
           Text(health)
            
            Button("get health/"){
                Task{
                   await getHealth()
                }
            }
        }
        .padding()
    }
    
    
    func getHealth() async {
        let url = URL(string:"http://localhost:6767/health")!
        var request = URLRequest(url: url )
        request.httpMethod = "GET"
        request.httpBody = nil
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        
        do {
            let (data, _) = try await URLSession.shared.data(for: request)
            health = String(data: data, encoding: .utf8) ?? "no data"
        } catch {
            
        }
        
        
    }

}

#Preview {
    ContentView()
}
