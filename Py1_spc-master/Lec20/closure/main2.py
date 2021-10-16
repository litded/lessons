def make_request(url, uri, postal_code, secure):
    """
    Функция, которая принимает много параметров, и все параметры - обязательные
    """
    print("Configure for url:", url)
    print("Configure for uri:", uri)
    print("Configure postal code:", postal_code)
    print("Check for secure:", secure)


make_request("vk.com", "mysql:///2233", "140041", "SUPER SECURE9000")
make_request("yandex.ru", "mysql:///2233", "140041", "SUPER SECURE9000")
make_request("google.com", "mysql:///2233", "140041", "SUPER SECURE9000")
make_request("habr.com", "mysql:///2233", "140041", "SUPER SECURE9000")
make_request("github.com", "mysql:///2233", "140041", "SUPER SECURE9000")

print("===================================================================")
def requester(uri, postal_code, secure):
    """
    Функция которая принимает наборо параметров, которые будут редко меняться
    """
    def new_request(url):
        print("Configure for url:", url)
        print("Configure for uri:", uri)
        print("Configure postal code:", postal_code)
        print("Check for secure:", secure)

    return new_request

current_request = requester("mysql:///2233", "140041", "SUPER SECURE9000")
current_request("vk.com")
current_request("yandex.ru")
current_request("google.com")
current_request("habr.com")
current_request("github.com")