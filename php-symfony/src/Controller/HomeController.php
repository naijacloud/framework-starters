<?php

namespace App\Controller;

use Symfony\Component\HttpFoundation\Response;
use Symfony\Component\Routing\Attribute\Route;

class HomeController
{
    #[Route('/', name: 'home')]
    public function index(): Response
    {
        return new Response('hello from symfony');
    }

    #[Route('/lucky', name: 'lucky')]
    public function lucky(): Response
    {
        return new Response('lucky route works');
    }
}
